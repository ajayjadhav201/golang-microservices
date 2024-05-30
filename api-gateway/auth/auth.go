package auth

import (
	"errors"
	"golang-microservices/api"
	"net/http"

	"golang-microservices/common"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/gin-gonic/gin"
)

type AuthClient struct {
	Client api.AuthServiceClient
	aws    *AwsS3Service
}

func NewAuthClient(service api.AuthServiceClient, aws *AwsS3Service) *AuthClient {
	return &AuthClient{service, aws}
}

func (a *AuthClient) RegisterRoutes(r *gin.RouterGroup) {
	// mux.HandleFunc("POST /api/v2/signup", a.SignupHandler)
	// mux.HandleFunc("POST /api/v2/uploadimage", a.uploadImage)
	// mux.HandleFunc("POST /api/v2/login", a.LoginHandler)
	// mux.HandleFunc("POST /api/v2/changepassword", a.ChangePassword)
	// mux.HandleFunc("POST /api/v2/updateuser", a.UpdateUserHandler)
	// mux.HandleFunc("POST /api/v2/deleteuser", a.DeleteUser)
}

//
//
//
//
//

// Signup
func (a *AuthClient) SignupHandler(w http.ResponseWriter, r *http.Request) {
	req := &api.SignupRequest{}
	if err := common.ReadJSON(r, req); err != nil {
		common.Println("ajaj error while parsing json", err.Error())
		common.WriteRequestBodyError(w, err)
		return
	}
	common.Println("ajaj signup request is ", req)

	res, err := a.Client.Signup(r.Context(), req)
	if err != nil {
		common.Println("ajaj signup errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

// Login
func (a *AuthClient) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req *api.LoginRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.Login(r.Context(), req)
	if err != nil {
		common.Println("ajaj login errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) ForgotPassword() {
	// mail sending
}

func (a *AuthClient) ChangePassword(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.ChangePasswordRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.ChangePassword(r.Context(), req)
	if err != nil {
		common.Println("ajaj chnage password errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.UpdateUserRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.UpdateUser(r.Context(), req)
	if err != nil {
		common.Println("ajaj updateuser errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

func (a *AuthClient) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//
	var req *api.DeleteUserRequest
	if err := common.ReadJSON(r, req); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}
	res, err := a.Client.DeleteUser(r.Context(), req)
	if err != nil {
		common.Println("ajaj deleteuser errror is ", err)
		common.WriteGrpcError(w, err)
		return
	}
	common.WriteJSON(w, http.StatusOK, res)
}

// upload profile image
func (a *AuthClient) uploadImage(w http.ResponseWriter, r *http.Request) {
	common.Println("request received as ", r.Body)
	if a.aws == nil {
		common.WriteError(w, http.StatusServiceUnavailable, "Service not available")
		return
	}
	err := r.ParseMultipartForm(2 << 20)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	form := r.MultipartForm
	if form == nil {
		common.WriteError(w, http.StatusBadRequest, "Unable to load form-data")
		return
	}
	//data := form.Value //these are form fileds

	files := form.File
	if len(files) == 0 {
		common.WriteError(w, http.StatusBadRequest, "Please upload files")
		return
	}
	common.Println("ajaj files are: ", files)

	var uploadedUrls []string

	for _, file := range files {
		fileHeaders := file
		if len(fileHeaders) == 0 {
			common.Println("ajaj fileheader is empty", fileHeaders)
			common.WriteError(w, http.StatusInternalServerError, "Unable to read Empty file")
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
				common.WriteError(w, http.StatusInternalServerError, common.Sprintf("Internal Server error: %s", errorid))
			}
			common.WriteError(w, http.StatusInternalServerError, common.Sprintf("Internal Server error: %s", err.Error()))
			return
		}
		uploadedUrls = append(uploadedUrls, url)
	}

	common.WriteJSON(w, http.StatusOK, common.Response{Message: common.Sprintf("Files uploaded successfully. path: %s", uploadedUrls[0])})
	return
}
