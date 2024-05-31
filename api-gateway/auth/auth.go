package auth

import (
	"errors"
	"golang-microservices/api/pb"
	"net/http"

	"golang-microservices/common"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gin-gonic/gin"
)

type AuthClient struct {
	Client pb.AuthServiceClient
	aws    *AwsS3Service
}

func NewAuthClient(service pb.AuthServiceClient, aws *AwsS3Service) *AuthClient {
	return &AuthClient{service, aws}
}

func (a *AuthClient) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/signup", a.SignupHandler)
	r.POST("/login", a.LoginHandler)
	r.POST("/uploadimage", AuthMiddleware, a.UploadImage)
	r.POST("/changepassword", a.ChangePassword)
	r.Use(AuthMiddleware)
	r.POST("/updateuser", a.UpdateUserHandler)
	r.POST("/deleteuser", a.DeleteUser)
}

//
//
//
//
//

// Signup
func (a *AuthClient) SignupHandler(c *gin.Context) {
	req := &pb.SignupRequest{}
	//
	if err := common.ReadJSON(c, req); err != nil {
		common.Println("ajaj error while parsing json", err.Error())
		common.WriteRequestBodyError(c, err)
		return
	}
	common.Println("ajaj signup request is ", req)
	//
	//
	res, err := a.Client.Signup(c.Request.Context(), req)
	if err != nil {
		common.Println("ajaj signup errror is ", err)
		common.WriteGrpcError(c, err)
		return
	}
	common.WriteJSON(c, http.StatusOK, res)
}

// Login
func (a *AuthClient) LoginHandler(c *gin.Context) {
	req := &pb.LoginRequest{}
	if err := common.ReadJSON(c, req); err != nil {
		common.WriteError(c, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.Login(c.Request.Context(), req)
	if err != nil {
		common.Println("ajaj login errror is ", err)
		common.WriteGrpcError(c, err)
		return
	}
	common.WriteJSON(c, http.StatusOK, res)
}

func (a *AuthClient) ForgotPassword() {
	// mail sending
}

func (a *AuthClient) ChangePassword(c *gin.Context) {

	var req *pb.ChangePasswordRequest
	if err := common.ReadJSON(c, req); err != nil {
		common.WriteError(c, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.ChangePassword(c.Request.Context(), req)
	if err != nil {
		common.Println("ajaj chnage password errror is ", err)
		common.WriteGrpcError(c, err)
		return
	}
	common.WriteJSON(c, http.StatusOK, res)
}

func (a *AuthClient) UpdateUserHandler(c *gin.Context) {
	// get userid from keys
	userid, ok := c.Get("userid")
	if !ok || userid == "" || common.Atoi(userid.(string), -1) == -1 {
		common.WriteError(c, http.StatusBadRequest, "update user error")
	}
	common.Println("ajaj request user id is: ", userid)
	//validate request
	req := &pb.UpdateUserRequest{}
	if err := common.ReadJSON(c, req); err != nil {
		common.WriteError(c, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	//requesting to user service
	req.UserID = userid.(string)
	res, err := a.Client.UpdateUser(c.Request.Context(), req)
	if err != nil {
		common.Println("ajaj updateuser errror is ", err)
		common.WriteGrpcError(c, err)
		return
	}
	common.WriteJSON(c, http.StatusOK, res)
}

func (a *AuthClient) DeleteUser(c *gin.Context) {
	//
	var req *pb.DeleteUserRequest
	if err := common.ReadJSON(c, req); err != nil {
		common.WriteError(c, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.DeleteUser(c.Request.Context(), req)
	if err != nil {
		common.Println("ajaj deleteuser errror is ", err)
		common.WriteGrpcError(c, err)
		return
	}
	common.WriteJSON(c, http.StatusOK, res)
}

// upload profile image
func (a *AuthClient) UploadImage(c *gin.Context) {
	common.Println("request received as ", c.Request.Body)
	if a.aws == nil {
		common.WriteError(c, http.StatusServiceUnavailable, "Service not available")
		return
	}
	err := c.Request.ParseMultipartForm(2 << 20)
	if err != nil {
		common.WriteError(c, http.StatusInternalServerError, err.Error())
		return
	}
	form := c.Request.MultipartForm
	if form == nil {
		common.WriteError(c, http.StatusBadRequest, "Unable to load form-data")
		return
	}
	//data := form.Value //these are form fileds

	files := form.File
	if len(files) == 0 {
		common.WriteError(c, http.StatusBadRequest, "Please upload files")
		return
	}
	common.Println("ajaj files are: ", files)

	var uploadedUrls []string

	for _, file := range files {
		fileHeaders := file
		if len(fileHeaders) == 0 {
			common.Println("ajaj fileheader is empty", fileHeaders)
			common.WriteError(c, http.StatusInternalServerError, "Unable to read Empty file")
			return
		}
		common.Println("ajaj fileheader is ", fileHeaders[0])
		url, err := UploadFile(a.aws.Uploader, a.aws.BucketName, fileHeaders[0])
		//
		if err != nil {
			common.Println(" error occured while uploading image", err)
			var mu manager.MultiUploadFailure
			if errors.As(err, &mu) {
				errorid := mu.UploadID() // retrieve the associated UploadID
				common.WriteError(c, http.StatusInternalServerError, common.Sprintf("Internal Server error: %s", errorid))
			}
			common.WriteError(c, http.StatusInternalServerError, common.Sprintf("Internal Server error: %s", err.Error()))
			return
		}
		uploadedUrls = append(uploadedUrls, url)
	}

	common.WriteJSON(c, http.StatusOK, common.Response{Message: common.Sprintf("Files uploaded successfully. path: %s", uploadedUrls[0])})
}
