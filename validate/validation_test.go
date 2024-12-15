package validate

import (
	"fmt"
	"testing"
	"time"
)


func TestValidatePhoneNumber(t *testing.T){
	//Valid phone number input. Should not throw error
	phonenumber1:="+2547856"
	if err:=ValidatePhoneNumber(phonenumber1); err != nil{
        t.Error(fmt.Sprintf("ValidatePhoneNumber(%s) is : %s", phonenumber1, ValidatePhoneNumber(phonenumber1)))
    }
	//Invalid phone number input. Should throw error
	phonenumber2:="2547856"
	if err:=ValidatePhoneNumber(phonenumber2); err == nil{
        t.Error(fmt.Sprintf("ValidatePhoneNumber(%s) is : %s", phonenumber2, ValidatePhoneNumber(phonenumber2)))
    }
    
	//Invalid phone number input. Should throw error
	phonenumber3:=""
	if err:=ValidatePhoneNumber(phonenumber3); err == nil{
        t.Error(fmt.Sprintf("ValidatePhoneNumber(%s) is : %s", phonenumber3, ValidatePhoneNumber(phonenumber3)))
    }
}

func TestValidateKenyanPhoneNumber(t *testing.T){
	//Valid Kenyan phone number input. Should not throw error
	phonenumber1:="+254756142241"
	if err:=ValidateKenyanPhoneNumber(phonenumber1); err != nil{
        t.Error(fmt.Sprintf("ValidateKenyanPhoneNumber(%s) is : %s", phonenumber1, ValidateKenyanPhoneNumber(phonenumber1)))
    }

	//Invalid Kenyan phone number input. Should throw error
	phonenumber2:="+2547856"
	if err:=ValidateKenyanPhoneNumber(phonenumber2); err == nil{
        t.Error(fmt.Sprintf("ValidateKenyanPhoneNumber(%s) is : %s", phonenumber2, ValidateKenyanPhoneNumber(phonenumber2)))
    }

	//Invalid Kenyan phone number input. Should throw error
	phonenumber3:=""
	if err:=ValidateKenyanPhoneNumber(phonenumber3); err == nil{
        t.Error(fmt.Sprintf("ValidateKenyanPhoneNumber(%s) is : %s", phonenumber3, ValidateKenyanPhoneNumber(phonenumber3)))
    }

	//Invalid Kenyan phone number input. Should throw error
	phonenumber4:="-254719226150"
	if err:=ValidateKenyanPhoneNumber(phonenumber4); err == nil{
        t.Error(fmt.Sprintf("ValidateKenyanPhoneNumber(%s) is : %s", phonenumber4, ValidateKenyanPhoneNumber(phonenumber4)))
    }
}
func TestValidateId(t *testing.T){

	//Valid id input. Should not throw error
	id1:= 6
	if err:=ValidateId(id1); err != nil{
        t.Error(fmt.Sprintf("ValidateId(%d) is : %s", id1, ValidateId(id1)))
    }

	//Invalid id input. Should throw error
	id2:= 0
	if err:=ValidateId(id2); err == nil{
        t.Error(fmt.Sprintf("ValidateId(%d) is : %s", id2, ValidateId(id2)))
    }

	//Invalid id input. Should throw error
	id3:=-3
	if err:=ValidateId(id3); err == nil{
        t.Error(fmt.Sprintf("ValidateId(%d) is : %s", id3, ValidateId(id3)))
    }
}

func TestValidateOtp(t *testing.T){

	//Valid OTP input. Should not throw error 
	Otp1:= "678453"
	if err:=ValidateOtp(Otp1); err != nil{
        t.Error(fmt.Sprintf("ValidateOtp(%s) is : %s", Otp1, ValidateOtp(Otp1)))
    }

	//Invalid OTP input. Should throw error 
	Otp2:= "785"
	if err:=ValidateOtp(Otp2); err == nil{
        t.Error(fmt.Sprintf("ValidateOtp(%s) is : %s", Otp2, ValidateOtp(Otp2)))
    }

	//Invalid OTP input. Should throw error 
	Otp3:="9845782939"
	if err:=ValidateOtp(Otp3); err == nil{
        t.Error(fmt.Sprintf("ValidateOtp(%s) is : %s", Otp3, ValidateOtp(Otp3)))
    }
}

func TestRegistrationNumber(t *testing.T){

	//Valid registration number input. Should not throw error 
	registration_number1:= "678453"
	if err:=ValidateRegistrationNumber(registration_number1); err != nil{
        t.Error(fmt.Sprintf("ValidateRegistrationNumber(%s) is : %s", registration_number1, ValidateRegistrationNumber(registration_number1)))
    }

	//Valid registration number input. Should not throw error 
	registration_number2:= "s"
	if err:=ValidateRegistrationNumber(registration_number2); err != nil{
        t.Error(fmt.Sprintf("ValidateRegistrationNumber(%s) is : %s", registration_number2, ValidateRegistrationNumber(registration_number2)))
    }

	//Valid registration number input. Should throw error 
	registration_number3:=""
	if err:=ValidateRegistrationNumber(registration_number3); err == nil{
        t.Error(fmt.Sprintf("ValidateRegistrationNumber(%s) is : %s", registration_number3, ValidateRegistrationNumber(registration_number3)))
    }
}  

func TestValidateName(t *testing.T){

	//Valid name input. Should not throw error 
	name1:= "Gwichu"
	if err:=ValidateName(name1); err != nil{
        t.Error(fmt.Sprintf("ValidateName(%s) is : %s", name1, ValidateName(name1)))
    }

	//Invalid name input. Should throw error 
	name2:= ""
	if err:=ValidateName(name2); err == nil{
        t.Error(fmt.Sprintf("ValidateName(%s) is : %s", name2, ValidateName(name2)))
    }
}  

func TestValidatePassword(t *testing.T){

	//Invalid  password input. Should throw error 
	password1:= "Gwichu"
	if err:=ValidatePassword(password1); err == nil{
        t.Error(fmt.Sprintf("ValidatePassword(%s) is : %s", password1, ValidatePassword(password1)))
    }

	//Invalid  password input. Should throw error 
	password2:= ""
	if err:=ValidatePassword(password2); err == nil{
        t.Error(fmt.Sprintf("ValidatePassword(%s) is : %s", password2, ValidatePassword(password2)))
    }

	//Valid  password input. Should not throw error 
	password3:= "vdcefkcnwcwefvnw"
	if err:=ValidatePassword(password3); err != nil{
        t.Error(fmt.Sprintf("ValidatePassword(%s) is : %s", password3, ValidatePassword(password3)))
    }
}  

func TestValidateDate(t *testing.T){

	//Invalid  date input. Should throw error 
	date1:= "2016"
	if err:=ValidateDate(date1); err == nil{
        t.Error(fmt.Sprintf("ValidateDate(%s) is : %s", date1, ValidateDate(date1)))
    }

	//Invalid  date input. Should throw error 
	date2:= ""
	if err:=ValidateDate(date2); err == nil{
        t.Error(fmt.Sprintf("ValidateDate(%s) is : %s", date2, ValidateDate(date2)))
    }

	//Valid date input. Should not throw error 
	date3:= "2016-03-02"
	if err:=ValidateDate(date3); err != nil{
        t.Error(fmt.Sprintf("ValidateDate(%s) is : %s", date3, ValidateDate(date3)))
    }
}  

func TestValidatePastDate(t *testing.T){
    
	//Valid  past date input. Should not throw error 
	date1:= time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC)

	if err:=ValidatePastDate(date1); err != nil{
        t.Error(fmt.Sprintf("ValidatePastDate(%s) is : %s", date1, ValidatePastDate(date1)))
    }

	//Valid  past date input. Should not throw error 
	date2:= time.Date(2024, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidatePastDate(date2); err != nil{
        t.Error(fmt.Sprintf("ValidatePastDate(%s) is : %s", date2, ValidatePastDate(date2)))
    }

	//Invalid  past date input. Should throw error 
	date3:= time.Date(2025, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidatePastDate(date3); err == nil{
        t.Error(fmt.Sprintf("ValidatePastDate(%s) is : %s", date3, ValidatePastDate(date3)))
    }
}  

func TestValidateDateOfAdmission(t *testing.T){

	//Valid  date of admission. Should not throw error 
	date1:= time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC)

	if err:=ValidateDateOfAdmission(date1); err != nil{
        t.Error(fmt.Sprintf("ValidateDateOfAdmission(%s) is : %s", date1, ValidateDateOfAdmission(date1)))
    }

	//Valid  date of admission. Should not throw error 
	date2:= time.Date(2024, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidatePastDate(date2); err != nil{
        t.Error(fmt.Sprintf("ValidateDateOfAdmission(%s) is : %s", date2, ValidateDateOfAdmission(date2)))
    }
    
	//Invalid  date of admission. Should throw error 
	date3:= time.Date(2025, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidatePastDate(date3); err == nil{
        t.Error(fmt.Sprintf("ValidateDateOfAdmission(%s) is : %s", date3, ValidateDateOfAdmission(date3)))
    }
}

func TestValidateDateOfBirth(t *testing.T){

	//Valid  date of birth. Should not throw error 
	date1:= time.Date(2023, 11, 23, 0, 0, 0, 0, time.UTC)

	if err:=ValidateDateOfBirth(date1); err != nil{
        t.Error(fmt.Sprintf("ValidateDateOfBirth(%s) is : %s", date1, ValidateDateOfBirth(date1)))
    }

	//Valid  date of birth. Should not throw error 
	date2:= time.Date(2024, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidateDateOfBirth(date2); err != nil{
        t.Error(fmt.Sprintf("ValidateDateOfBirth(%s) is : %s", date2, ValidateDateOfBirth(date2)))
    }

	//Invalid  date of birth. Should throw error 
	date3:= time.Date(2025, 12, 14, 0, 0, 0, 0, time.UTC)

	if err:=ValidateDateOfBirth(date3); err == nil{
        t.Error(fmt.Sprintf("ValidateDateOfBirth(%s) is : %s", date3, ValidateDateOfBirth(date3)))
    }
}

func TestValidateUrl(t *testing.T){

	//Valid  url. Should not throw error 
	url1:= "https://console-preview.neo4j.io/tools/query"

	if err:=ValidateUrl(url1); err != nil{
        t.Error(fmt.Sprintf("ValidateUrl(%s) is : %s", url1, ValidateUrl(url1)))
    }

	//Valid  url. Should not throw error 
	url2:= "http://console-preview.neo4j.io/tools/query"

	if err:=ValidateUrl(url2); err != nil{
        t.Error(fmt.Sprintf("ValidateUrl(%s) is : %s", url2, ValidateUrl(url2)))
    }

	//Invalid  url. Should throw error 
	url3:= "://console-preview.neo4j.io/tools/query"

	if err:=ValidateUrl(url3); err == nil{
        t.Error(fmt.Sprintf("ValidateUrl(%s) is : %s", url3, ValidateUrl(url3)))
    }
}  

func TestValidateProfilePicture(t *testing.T){

	//Valid  profile picture url. Should not throw error 
	url1:= "https://console-preview.neo4j.io/tools/query"

	if err:=ValidateProfilePicture(url1); err != nil{
        t.Error(fmt.Sprintf("ValidateProfilePicture(%s) is : %s", url1, ValidateProfilePicture(url1)))
    }

	//Valid  profile picture url. Should not throw error 
	url2:= "http://console-preview.neo4j.io/tools/query"

	if err:=ValidateProfilePicture(url2); err != nil{
        t.Error(fmt.Sprintf("ValidateProfilePicture(%s) is : %s", url2, ValidateProfilePicture(url2)))
    }

	//Invalid profile picture url. Should throw error 
	url3:= "://console-preview.neo4j.io/tools/query"

	if err:=ValidateUrl(url3); err == nil{
        t.Error(fmt.Sprintf("ValidateProfilePicture(%s) is : %s", url3, ValidateProfilePicture(url3)))
    }
} 

func TestValidateBadge(t *testing.T){

	//Valid  badge picture url. Should not throw error 
	url1:= "https://console-preview.neo4j.io/tools/query"

	if err:=ValidateBadge(url1); err != nil{
        t.Error(fmt.Sprintf("ValidateBadge(%s) is : %s", url1, ValidateBadge(url1)))
    }

	//Valid  badge picture url. Should not throw error 
	url2:= "http://console-preview.neo4j.io/tools/query"

	if err:=ValidateBadge(url2); err != nil{
        t.Error(fmt.Sprintf("ValidateBadge(%s) is : %s", url2, ValidateBadge(url2)))
    }

	//Invalid  badge picture url. Should throw error 
	url3:= "h://console-preview.neo4j.io/tools/query"

	if err:=ValidateBadge(url3); err == nil{
        t.Error(fmt.Sprintf("ValidateBadge(%s) is : %s", url3, ValidateBadge(url3)))
    }
}  

func TestValidateWebsite(t *testing.T){

	//Valid  website url. Should not throw error 
	url1:= "https://console-preview.neo4j.io/tools/query"

	if err:=ValidateWebsite(url1); err != nil{
        t.Error(fmt.Sprintf("ValidateWebsite(%s) is : %s", url1, ValidateWebsite(url1)))
    }

	//Valid  website url. Should not throw error 
	url2:= "http://console-preview.neo4j.io/tools/query"

	if err:=ValidateWebsite(url2); err != nil{
        t.Error(fmt.Sprintf("ValidateWebsite(%s) is : %s", url2, ValidateWebsite(url2)))
    }

	//Invalid website url. Should throw error 
	url3:= "h://console-preview.neo4j.io/tools/query"

	if err:=ValidateWebsite(url3); err == nil{
        t.Error(fmt.Sprintf("ValidateWebsite(%s) is : %s", url3, ValidateWebsite(url3)))
    }
}  