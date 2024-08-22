interface Kols {
	KolID: number;
	UserProfileID: number;
	Language: string;
	Education: string;
	ExpectedSalary: number;
	ExpectedSalaryEnable: boolean;
	ChannelSettingTypeID: number;
	IDFrontURL: string;
	IDBackURL: string;
	PortraitURL: string;
	RewardID: number;
	PaymentMethodID: number;
	TestimonialsID: number;
	VerificationStatus: boolean;
	Enabled: boolean;
	ActiveDate: Date;
	Active: boolean;
	CreatedBy: string;
	CreatedDate: Date;
	ModifiedBy: string;
	ModifiedDate: Date;
	IsRemove: boolean;
	IsOnBoarding: boolean;
	Code: string;
	PortraitRightURL: string;
	PortraitLeftURL: string;
	LivenessStatus: boolean;
}

export default Kols;