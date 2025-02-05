// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type AccountTakeoverEventActionType string

const (
	AccountTakeoverEventActionType_BLOCK             AccountTakeoverEventActionType = "BLOCK"
	AccountTakeoverEventActionType_MFA_IF_CONFIGURED AccountTakeoverEventActionType = "MFA_IF_CONFIGURED"
	AccountTakeoverEventActionType_MFA_REQUIRED      AccountTakeoverEventActionType = "MFA_REQUIRED"
	AccountTakeoverEventActionType_NO_ACTION         AccountTakeoverEventActionType = "NO_ACTION"
)

type AdvancedSecurityEnabledModeType string

const (
	AdvancedSecurityEnabledModeType_AUDIT    AdvancedSecurityEnabledModeType = "AUDIT"
	AdvancedSecurityEnabledModeType_ENFORCED AdvancedSecurityEnabledModeType = "ENFORCED"
)

type AdvancedSecurityModeType string

const (
	AdvancedSecurityModeType_AUDIT    AdvancedSecurityModeType = "AUDIT"
	AdvancedSecurityModeType_ENFORCED AdvancedSecurityModeType = "ENFORCED"
	AdvancedSecurityModeType_OFF      AdvancedSecurityModeType = "OFF"
)

type AliasAttributeType string

const (
	AliasAttributeType_email              AliasAttributeType = "email"
	AliasAttributeType_phone_number       AliasAttributeType = "phone_number"
	AliasAttributeType_preferred_username AliasAttributeType = "preferred_username"
)

type AssetCategoryType string

const (
	AssetCategoryType_AUTH_APP_GRAPHIC       AssetCategoryType = "AUTH_APP_GRAPHIC"
	AssetCategoryType_EMAIL_GRAPHIC          AssetCategoryType = "EMAIL_GRAPHIC"
	AssetCategoryType_FAVICON_ICO            AssetCategoryType = "FAVICON_ICO"
	AssetCategoryType_FAVICON_SVG            AssetCategoryType = "FAVICON_SVG"
	AssetCategoryType_FORM_BACKGROUND        AssetCategoryType = "FORM_BACKGROUND"
	AssetCategoryType_FORM_LOGO              AssetCategoryType = "FORM_LOGO"
	AssetCategoryType_IDP_BUTTON_ICON        AssetCategoryType = "IDP_BUTTON_ICON"
	AssetCategoryType_PAGE_BACKGROUND        AssetCategoryType = "PAGE_BACKGROUND"
	AssetCategoryType_PAGE_FOOTER_BACKGROUND AssetCategoryType = "PAGE_FOOTER_BACKGROUND"
	AssetCategoryType_PAGE_FOOTER_LOGO       AssetCategoryType = "PAGE_FOOTER_LOGO"
	AssetCategoryType_PAGE_HEADER_BACKGROUND AssetCategoryType = "PAGE_HEADER_BACKGROUND"
	AssetCategoryType_PAGE_HEADER_LOGO       AssetCategoryType = "PAGE_HEADER_LOGO"
	AssetCategoryType_PASSKEY_GRAPHIC        AssetCategoryType = "PASSKEY_GRAPHIC"
	AssetCategoryType_PASSWORD_GRAPHIC       AssetCategoryType = "PASSWORD_GRAPHIC"
	AssetCategoryType_SMS_GRAPHIC            AssetCategoryType = "SMS_GRAPHIC"
)

type AssetExtensionType string

const (
	AssetExtensionType_ICO  AssetExtensionType = "ICO"
	AssetExtensionType_JPEG AssetExtensionType = "JPEG"
	AssetExtensionType_PNG  AssetExtensionType = "PNG"
	AssetExtensionType_SVG  AssetExtensionType = "SVG"
	AssetExtensionType_WEBP AssetExtensionType = "WEBP"
)

type AttributeDataType string

const (
	AttributeDataType_Boolean  AttributeDataType = "Boolean"
	AttributeDataType_DateTime AttributeDataType = "DateTime"
	AttributeDataType_Number   AttributeDataType = "Number"
	AttributeDataType_String   AttributeDataType = "String"
)

type AuthFactorType string

const (
	AuthFactorType_EMAIL_OTP AuthFactorType = "EMAIL_OTP"
	AuthFactorType_PASSWORD  AuthFactorType = "PASSWORD"
	AuthFactorType_SMS_OTP   AuthFactorType = "SMS_OTP"
	AuthFactorType_WEB_AUTHN AuthFactorType = "WEB_AUTHN"
)

type AuthFlowType string

const (
	AuthFlowType_ADMIN_NO_SRP_AUTH        AuthFlowType = "ADMIN_NO_SRP_AUTH"
	AuthFlowType_ADMIN_USER_PASSWORD_AUTH AuthFlowType = "ADMIN_USER_PASSWORD_AUTH"
	AuthFlowType_CUSTOM_AUTH              AuthFlowType = "CUSTOM_AUTH"
	AuthFlowType_REFRESH_TOKEN            AuthFlowType = "REFRESH_TOKEN"
	AuthFlowType_REFRESH_TOKEN_AUTH       AuthFlowType = "REFRESH_TOKEN_AUTH"
	AuthFlowType_USER_AUTH                AuthFlowType = "USER_AUTH"
	AuthFlowType_USER_PASSWORD_AUTH       AuthFlowType = "USER_PASSWORD_AUTH"
	AuthFlowType_USER_SRP_AUTH            AuthFlowType = "USER_SRP_AUTH"
)

type ChallengeName string

const (
	ChallengeName_Mfa      ChallengeName = "Mfa"
	ChallengeName_Password ChallengeName = "Password"
)

type ChallengeNameType string

const (
	ChallengeNameType_ADMIN_NO_SRP_AUTH        ChallengeNameType = "ADMIN_NO_SRP_AUTH"
	ChallengeNameType_CUSTOM_CHALLENGE         ChallengeNameType = "CUSTOM_CHALLENGE"
	ChallengeNameType_DEVICE_PASSWORD_VERIFIER ChallengeNameType = "DEVICE_PASSWORD_VERIFIER"
	ChallengeNameType_DEVICE_SRP_AUTH          ChallengeNameType = "DEVICE_SRP_AUTH"
	ChallengeNameType_EMAIL_OTP                ChallengeNameType = "EMAIL_OTP"
	ChallengeNameType_MFA_SETUP                ChallengeNameType = "MFA_SETUP"
	ChallengeNameType_NEW_PASSWORD_REQUIRED    ChallengeNameType = "NEW_PASSWORD_REQUIRED"
	ChallengeNameType_PASSWORD                 ChallengeNameType = "PASSWORD"
	ChallengeNameType_PASSWORD_SRP             ChallengeNameType = "PASSWORD_SRP"
	ChallengeNameType_PASSWORD_VERIFIER        ChallengeNameType = "PASSWORD_VERIFIER"
	ChallengeNameType_SELECT_CHALLENGE         ChallengeNameType = "SELECT_CHALLENGE"
	ChallengeNameType_SELECT_MFA_TYPE          ChallengeNameType = "SELECT_MFA_TYPE"
	ChallengeNameType_SMS_MFA                  ChallengeNameType = "SMS_MFA"
	ChallengeNameType_SMS_OTP                  ChallengeNameType = "SMS_OTP"
	ChallengeNameType_SOFTWARE_TOKEN_MFA       ChallengeNameType = "SOFTWARE_TOKEN_MFA"
	ChallengeNameType_WEB_AUTHN                ChallengeNameType = "WEB_AUTHN"
)

type ChallengeResponse string

const (
	ChallengeResponse_Failure ChallengeResponse = "Failure"
	ChallengeResponse_Success ChallengeResponse = "Success"
)

type ColorSchemeModeType string

const (
	ColorSchemeModeType_DARK    ColorSchemeModeType = "DARK"
	ColorSchemeModeType_DYNAMIC ColorSchemeModeType = "DYNAMIC"
	ColorSchemeModeType_LIGHT   ColorSchemeModeType = "LIGHT"
)

type CompromisedCredentialsEventActionType string

const (
	CompromisedCredentialsEventActionType_BLOCK     CompromisedCredentialsEventActionType = "BLOCK"
	CompromisedCredentialsEventActionType_NO_ACTION CompromisedCredentialsEventActionType = "NO_ACTION"
)

type CustomEmailSenderLambdaVersionType string

const (
	CustomEmailSenderLambdaVersionType_V1_0 CustomEmailSenderLambdaVersionType = "V1_0"
)

type CustomSMSSenderLambdaVersionType string

const (
	CustomSMSSenderLambdaVersionType_V1_0 CustomSMSSenderLambdaVersionType = "V1_0"
)

type DefaultEmailOptionType string

const (
	DefaultEmailOptionType_CONFIRM_WITH_CODE DefaultEmailOptionType = "CONFIRM_WITH_CODE"
	DefaultEmailOptionType_CONFIRM_WITH_LINK DefaultEmailOptionType = "CONFIRM_WITH_LINK"
)

type DeletionProtectionType string

const (
	DeletionProtectionType_ACTIVE   DeletionProtectionType = "ACTIVE"
	DeletionProtectionType_INACTIVE DeletionProtectionType = "INACTIVE"
)

type DeliveryMediumType string

const (
	DeliveryMediumType_EMAIL DeliveryMediumType = "EMAIL"
	DeliveryMediumType_SMS   DeliveryMediumType = "SMS"
)

type DeviceRememberedStatusType string

const (
	DeviceRememberedStatusType_not_remembered DeviceRememberedStatusType = "not_remembered"
	DeviceRememberedStatusType_remembered     DeviceRememberedStatusType = "remembered"
)

type DomainStatusType string

const (
	DomainStatusType_ACTIVE   DomainStatusType = "ACTIVE"
	DomainStatusType_CREATING DomainStatusType = "CREATING"
	DomainStatusType_DELETING DomainStatusType = "DELETING"
	DomainStatusType_FAILED   DomainStatusType = "FAILED"
	DomainStatusType_UPDATING DomainStatusType = "UPDATING"
)

type EmailSendingAccountType string

const (
	EmailSendingAccountType_COGNITO_DEFAULT EmailSendingAccountType = "COGNITO_DEFAULT"
	EmailSendingAccountType_DEVELOPER       EmailSendingAccountType = "DEVELOPER"
)

type EventFilterType string

const (
	EventFilterType_PASSWORD_CHANGE EventFilterType = "PASSWORD_CHANGE"
	EventFilterType_SIGN_IN         EventFilterType = "SIGN_IN"
	EventFilterType_SIGN_UP         EventFilterType = "SIGN_UP"
)

type EventResponseType string

const (
	EventResponseType_Fail       EventResponseType = "Fail"
	EventResponseType_InProgress EventResponseType = "InProgress"
	EventResponseType_Pass       EventResponseType = "Pass"
)

type EventSourceName string

const (
	EventSourceName_userAuthEvents   EventSourceName = "userAuthEvents"
	EventSourceName_userNotification EventSourceName = "userNotification"
)

type EventType string

const (
	EventType_ForgotPassword EventType = "ForgotPassword"
	EventType_PasswordChange EventType = "PasswordChange"
	EventType_ResendCode     EventType = "ResendCode"
	EventType_SignIn         EventType = "SignIn"
	EventType_SignUp         EventType = "SignUp"
)

type ExplicitAuthFlowsType string

const (
	ExplicitAuthFlowsType_ADMIN_NO_SRP_AUTH              ExplicitAuthFlowsType = "ADMIN_NO_SRP_AUTH"
	ExplicitAuthFlowsType_ALLOW_ADMIN_USER_PASSWORD_AUTH ExplicitAuthFlowsType = "ALLOW_ADMIN_USER_PASSWORD_AUTH"
	ExplicitAuthFlowsType_ALLOW_CUSTOM_AUTH              ExplicitAuthFlowsType = "ALLOW_CUSTOM_AUTH"
	ExplicitAuthFlowsType_ALLOW_REFRESH_TOKEN_AUTH       ExplicitAuthFlowsType = "ALLOW_REFRESH_TOKEN_AUTH"
	ExplicitAuthFlowsType_ALLOW_USER_AUTH                ExplicitAuthFlowsType = "ALLOW_USER_AUTH"
	ExplicitAuthFlowsType_ALLOW_USER_PASSWORD_AUTH       ExplicitAuthFlowsType = "ALLOW_USER_PASSWORD_AUTH"
	ExplicitAuthFlowsType_ALLOW_USER_SRP_AUTH            ExplicitAuthFlowsType = "ALLOW_USER_SRP_AUTH"
	ExplicitAuthFlowsType_CUSTOM_AUTH_FLOW_ONLY          ExplicitAuthFlowsType = "CUSTOM_AUTH_FLOW_ONLY"
	ExplicitAuthFlowsType_USER_PASSWORD_AUTH             ExplicitAuthFlowsType = "USER_PASSWORD_AUTH"
)

type FeedbackValueType string

const (
	FeedbackValueType_Invalid FeedbackValueType = "Invalid"
	FeedbackValueType_Valid   FeedbackValueType = "Valid"
)

type IdentityProviderTypeType string

const (
	IdentityProviderTypeType_Facebook        IdentityProviderTypeType = "Facebook"
	IdentityProviderTypeType_Google          IdentityProviderTypeType = "Google"
	IdentityProviderTypeType_LoginWithAmazon IdentityProviderTypeType = "LoginWithAmazon"
	IdentityProviderTypeType_OIDC            IdentityProviderTypeType = "OIDC"
	IdentityProviderTypeType_SAML            IdentityProviderTypeType = "SAML"
	IdentityProviderTypeType_SignInWithApple IdentityProviderTypeType = "SignInWithApple"
)

type LogLevel string

const (
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_INFO  LogLevel = "INFO"
)

type MessageActionType string

const (
	MessageActionType_RESEND   MessageActionType = "RESEND"
	MessageActionType_SUPPRESS MessageActionType = "SUPPRESS"
)

type OAuthFlowType string

const (
	OAuthFlowType_client_credentials OAuthFlowType = "client_credentials"
	OAuthFlowType_code               OAuthFlowType = "code"
	OAuthFlowType_implicit           OAuthFlowType = "implicit"
)

type PreTokenGenerationLambdaVersionType string

const (
	PreTokenGenerationLambdaVersionType_V1_0 PreTokenGenerationLambdaVersionType = "V1_0"
	PreTokenGenerationLambdaVersionType_V2_0 PreTokenGenerationLambdaVersionType = "V2_0"
)

type PreventUserExistenceErrorTypes string

const (
	PreventUserExistenceErrorTypes_ENABLED PreventUserExistenceErrorTypes = "ENABLED"
	PreventUserExistenceErrorTypes_LEGACY  PreventUserExistenceErrorTypes = "LEGACY"
)

type RecoveryOptionNameType string

const (
	RecoveryOptionNameType_admin_only            RecoveryOptionNameType = "admin_only"
	RecoveryOptionNameType_verified_email        RecoveryOptionNameType = "verified_email"
	RecoveryOptionNameType_verified_phone_number RecoveryOptionNameType = "verified_phone_number"
)

type RiskDecisionType string

const (
	RiskDecisionType_AccountTakeover RiskDecisionType = "AccountTakeover"
	RiskDecisionType_Block           RiskDecisionType = "Block"
	RiskDecisionType_NoRisk          RiskDecisionType = "NoRisk"
)

type RiskLevelType string

const (
	RiskLevelType_High   RiskLevelType = "High"
	RiskLevelType_Low    RiskLevelType = "Low"
	RiskLevelType_Medium RiskLevelType = "Medium"
)

type StatusType string

const (
	StatusType_Disabled StatusType = "Disabled"
	StatusType_Enabled  StatusType = "Enabled"
)

type TimeUnitsType string

const (
	TimeUnitsType_days    TimeUnitsType = "days"
	TimeUnitsType_hours   TimeUnitsType = "hours"
	TimeUnitsType_minutes TimeUnitsType = "minutes"
	TimeUnitsType_seconds TimeUnitsType = "seconds"
)

type UserImportJobStatusType string

const (
	UserImportJobStatusType_Created    UserImportJobStatusType = "Created"
	UserImportJobStatusType_Expired    UserImportJobStatusType = "Expired"
	UserImportJobStatusType_Failed     UserImportJobStatusType = "Failed"
	UserImportJobStatusType_InProgress UserImportJobStatusType = "InProgress"
	UserImportJobStatusType_Pending    UserImportJobStatusType = "Pending"
	UserImportJobStatusType_Stopped    UserImportJobStatusType = "Stopped"
	UserImportJobStatusType_Stopping   UserImportJobStatusType = "Stopping"
	UserImportJobStatusType_Succeeded  UserImportJobStatusType = "Succeeded"
)

type UserPoolMFAType string

const (
	UserPoolMFAType_OFF      UserPoolMFAType = "OFF"
	UserPoolMFAType_ON       UserPoolMFAType = "ON"
	UserPoolMFAType_OPTIONAL UserPoolMFAType = "OPTIONAL"
)

type UserPoolTierType string

const (
	UserPoolTierType_ESSENTIALS UserPoolTierType = "ESSENTIALS"
	UserPoolTierType_LITE       UserPoolTierType = "LITE"
	UserPoolTierType_PLUS       UserPoolTierType = "PLUS"
)

type UserStatusType string

const (
	UserStatusType_ARCHIVED              UserStatusType = "ARCHIVED"
	UserStatusType_COMPROMISED           UserStatusType = "COMPROMISED"
	UserStatusType_CONFIRMED             UserStatusType = "CONFIRMED"
	UserStatusType_EXTERNAL_PROVIDER     UserStatusType = "EXTERNAL_PROVIDER"
	UserStatusType_FORCE_CHANGE_PASSWORD UserStatusType = "FORCE_CHANGE_PASSWORD"
	UserStatusType_RESET_REQUIRED        UserStatusType = "RESET_REQUIRED"
	UserStatusType_UNCONFIRMED           UserStatusType = "UNCONFIRMED"
	UserStatusType_UNKNOWN               UserStatusType = "UNKNOWN"
)

type UserVerificationType string

const (
	UserVerificationType_preferred UserVerificationType = "preferred"
	UserVerificationType_required  UserVerificationType = "required"
)

type UsernameAttributeType string

const (
	UsernameAttributeType_email        UsernameAttributeType = "email"
	UsernameAttributeType_phone_number UsernameAttributeType = "phone_number"
)

type VerifiedAttributeType string

const (
	VerifiedAttributeType_email        VerifiedAttributeType = "email"
	VerifiedAttributeType_phone_number VerifiedAttributeType = "phone_number"
)

type VerifySoftwareTokenResponseType string

const (
	VerifySoftwareTokenResponseType_ERROR   VerifySoftwareTokenResponseType = "ERROR"
	VerifySoftwareTokenResponseType_SUCCESS VerifySoftwareTokenResponseType = "SUCCESS"
)
