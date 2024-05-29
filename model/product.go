package model

import "time"

type Product struct {
	ID         int
	Name       string
	Brand      string
	Category   string
	Images     []string // jsom marshaled list of string or pq.StringArray
	Quantity   int
	ComingSoon bool
	OutOfStock bool //currently unavailable
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Reviews    string
	//
	Price           float64
	Discount        int
	DiscountPercent float32
	FinalPrice      float64
	//
	//
	BattaryCapacity int // 5000 mAh
	BattaryType     string
	//
	//
	Rating           float32 // will not be shown
	TotalRatingCount int     // will not be shown
	//
	//
	RAM             int
	InternalStorage int
	Expandable      string //upto 1 TB
	//
	//
	ScreenSize      float32 // 6.97
	DisplaySize     string  // 17.25 cm (6.79 inch)
	DisplayType     string  // FHD+ 90Hz AdaptiveSync Display
	TouchScreenType string  // Capacitive
	ResolutionType  string  // Full HD+
	RefreshRate     int     // Hz
	UserInterface   string  //MIUI 14
	//
	//
	PrimaryCamera           int
	FrontCamera             string
	PrimaryCameraFeatures   string
	SecondaryCamera         int
	RearCamera              string
	SecondaryCameraFeatures string
	Flash                   string
	HDRecording             bool
	FullHDRecording         bool
	RecordingResolution     string
	ZoomType                string
	ZoomUpto                string
	FrameRate               string
	//
	//General
	InTheBox      string
	ModelNumber   string
	ModelName     string
	Color         string
	SimType       string // dual sim
	HybridSimSlot bool
	OTGCompatible bool
	QuickCharging bool
	SARValue      string
	//
	//
	OperatingSystem        string
	OperatingSystemVersion string
	ProcessorBrand         string  // Snapdragon
	ProcessorType          string  // Snapdragon 4 Gen 2
	ClockSpeed             float32 //primary
	SecondaryClockSpeed    float32
	//
	//
	NetworkType      string // 5g or 4G
	BluetoothSupport bool
	BluetoothVersion string
	Wifi             bool
	WifiVersion      string
	NFC              bool
	USBConnectivity  bool
	AudioJack        bool
	GPS              bool
	//
	//
	Sensors       string
	Browser       string
	OtherFeatures string
	FMRadio       bool
	//
	// Dimentions
	Width  float32
	Height float32
	Depth  float32 //in mm
	Weight int     //in grams
	//
	// Warranty
	WarrantySummery     string
	WarrentyPeriod      string
	WarrantyServicetype string //NA
	//
	//
	CountryOfOrigin     string
	ManufacturerDetails string
	ImporterDetails     string
	PackerDetails       string
}

/*
Color
Storage  ram
Name/Title
rating
totoalRatingCount
price
discount
sellingPrice
Quick Summery []string
//
Processor
Rear Camera
Front camera
Display
otherDetails map[string]string
	Network Type		string
	Sim Type
	Expandable Storage 	bool
	Audio Jack			bool
	Quick Charging 		bool
	in the box 			string


Description map[string]string
	imagePath			string
	description			string

// Specifications
Key_Features []string

General [string]string
	in the box			string
	Model_number		string
	Model_Name			string
	Color				string
	Sim_type			string
	Hybrid sim slot		bool
	TouchScreen			string
	OTG compatible		bool

Display Features [string]string
	Display_Size		string
	resolution			string
	resolution_type		string
	Display_type		string
	other disp.Features	string

os &Processor Features [string]string
	operating system	string
	Processor Brand		string  snapdragon
	Processor Type		string	Snapdragon 4 Gen 2
	Processor core		string	Octa Core
	Primary clockspeed	string	2.2 GHz
	Secondary c.Speed	string	2 GHz
	operating Frequency	string

Memory and Storage Features	[string]string
	Internal_storage	string
	RAm					string

Camera Features [string]string
	primary camera		string
	primary c. features	string
	secnd.camera		string
	sec. cam.Features	string
	Flash				string
	Full HD Recording	bool

Connectivity Features [string]string
	Network Type		string
	Supported Networks	string
	Internal Connectivity string
	Micor USb version	string
	Bluetooth Support	bool
	Bluetooth version	string
	wifi				string
	Wifi version		string
	Wifi-Hotspot		bool
	NFC 				bool
	USB Connectivity	bool
	GPS Support			bool

Other Details [string]string
	smartPhone			string
	TouchScreen type	string
	Sim Size			string
	User Interface		string
	Instant Message		bool
	SMS					bool
	Graphics PPi		string
	Sensors				string
	Other Features		string
	GPS type			string
	FM Radio			bool
	Battary Capacity 	string
	Battary TYpe		string


Dimentions	[string]string
	Width				string
	Height				string
	Depth				string
	Weight				string


More_info map[string]string
	Warranty			string
	Mfg. details		string
	importer details	string
	packer details		string



*/

/*
Filters
	price
		Rs. 10000 and Below
		Rs. 10000 - Rs. 15000
		Rs. 15000 - Rs. 20000
		Rs. 20000 - Rs. 30000
		Rs. 30000 and Above

	Brand
		samsung
		vivo
		APPLE
		Google
		Redmi

	Ratings
		4* and Above
		3* and above

	RAM
		4 GB
		6 GB
		8 GB and Above

	Internal Storage
		32 GB
		64 GB
		128 GB
		256 GB
		512 GB
		1 TB

	Battary Capacity
		3000 - 3999 mAh
		4000 - 4999 mAh
		5000 - 5999 mAh
		6000 mAh and Above

	Screen Size
		5.0 - 5.9 inch
		6 inch and Above

	Primary Camera
		21 MP and Below
		21 MP and Above
		50 MP and Above
		64 Mp and Above

	Secondary Camera
		21 MP and Below
		21 MP and Above
		50 MP and Above
		64 Mp and Above

	Processor Brand
		Exynos
		Mediatek
		Snapdragon

	Resolution type
		HD
		Full HD
		Full HD+
		Quad HD
		Quad HD+
		Full HD+ AMOLED Display
		Full HD+ Super AMOLED Display
		Liquid Retina Display
		UHD 4k

	Operating System
		IOS
		Android

	Network Type
		2G
		3G
		4G
		4G VOLTE
		5G

	Availability
		Exclude Out of Stock  bool

	Discount
		10% and more
		20% and more
		30% and more
		40% and more
		50% and more

	Clock Speed
		1.5 - 1.9 GHz
		2GHz and Above
*/
