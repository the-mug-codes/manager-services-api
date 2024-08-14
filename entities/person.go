package entities

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/the-mug-codes/adapters-service-api/database/utils"
	"gorm.io/gorm"
)

type mainPhone struct {
	Phone     PersonPhone
	Formatted string
	AreaPhone string
	FullPhone string
}

type PersonPhone struct {
	ID          uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	PersonID    uuid.UUID       `json:"person_id" binding:"required"`
	CountryCode int             `json:"country_code" binding:"required"`
	AreaCode    int             `json:"area_code" binding:"required"`
	PhoneNumber int             `json:"phone_number" binding:"required"`
	IsWhatsApp  bool            `json:"is_whatsapp"`
	IsMain      bool            `json:"is_main"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PersonEmail struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	PersonID  uuid.UUID       `json:"person_id" binding:"required"`
	Email     string          `json:"email"`
	IsMain    bool            `json:"is_main"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PersonAddress struct {
	ID           uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	PersonID     string          `json:"person_id" binding:"required"`
	Address      string          `json:"address" binding:"required"`
	Number       *string         `json:"number,omitempty"`
	Complement   *string         `json:"complement,omitempty"`
	City         string          `json:"city" binding:"required"`
	State        string          `json:"state" binding:"required"`
	ZipCode      string          `json:"zip_code" binding:"required"`
	Neighborhood string          `json:"neighborhood" binding:"required"`
	Country      string          `json:"country" binding:"required"`
	IsMain       bool            `json:"is_main"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PersonExtra struct {
	ID        uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	PersonID  string          `json:"person_id" binding:"required"`
	Key       string          `json:"key" binding:"required"`
	Value     string          `json:"value" binding:"required"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type Person struct {
	ID           uuid.UUID       `gorm:"type:uuid;primary_key" json:"id"`
	Nickname     *string         `json:"nickname"`
	Name         string          `json:"name" binding:"required"`
	Document     string          `gorm:"index;unique" json:"document" binding:"required"`
	MainPersonID *uuid.UUID      `json:"main_person_id"`
	Addresses    []PersonAddress `gorm:"foreignKey:PersonID;constraint:OnDelete:CASCADE" json:"addresses"`
	Emails       []PersonEmail   `gorm:"foreignKey:PersonID;constraint:OnDelete:CASCADE" json:"emails"`
	Phones       []PersonPhone   `gorm:"foreignKey:PersonID;constraint:OnDelete:CASCADE" json:"phones"`
	Extras       []PersonExtra   `gorm:"foreignKey:PersonID;constraint:OnDelete:CASCADE" json:"extras"`
	Status       bool            `gorm:"default:true" json:"status"`
	CreatedAt    time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (phone *PersonPhone) formattedPhone() string {
	pattern := `^(\d{4,5})(\d{4})$`
	mask := "$1-$2"
	re := regexp.MustCompile(pattern)
	phoneNumber := re.ReplaceAllString(fmt.Sprintf("%v", phone.PhoneNumber), mask)
	return fmt.Sprintf("+%v (%v) %v", phone.CountryCode, phone.AreaCode, phoneNumber)
}

func (phone *PersonPhone) fullPhone(country bool) string {
	if country {
		return fmt.Sprintf("%v%v%v", phone.CountryCode, phone.AreaCode, phone.PhoneNumber)
	}
	return fmt.Sprintf("%v%v", phone.AreaCode, phone.PhoneNumber)
}

func (person *Person) GetMainPhone(whatsapp bool) (main mainPhone) {
	if whatsapp {
		for _, phone := range person.Phones {
			if phone.IsMain && phone.IsWhatsApp {
				main = mainPhone{
					Phone:     phone,
					Formatted: phone.formattedPhone(),
					AreaPhone: phone.fullPhone(false),
					FullPhone: phone.fullPhone(true),
				}
			}
		}
		return main
	}
	for _, phone := range person.Phones {
		if phone.IsMain {
			main = mainPhone{
				Phone:     phone,
				Formatted: phone.formattedPhone(),
				AreaPhone: phone.fullPhone(false),
				FullPhone: phone.fullPhone(true),
			}
		}
	}
	return main
}

func (person *Person) GetMainEmail() (main PersonEmail) {
	for _, email := range person.Emails {
		if email.IsMain {
			return email
		}
	}
	return main
}

func (person Person) GetMainAddress() string {
	for _, address := range person.Addresses {
		if address.IsMain {
			var newAddress strings.Builder
			pattern := `^(\d{5})(\d{3})$`
			mask := "$1-$2"
			re := regexp.MustCompile(pattern)
			zipCode := re.ReplaceAllString(address.ZipCode, mask)
			newAddress.WriteString(fmt.Sprintf("%s, ", address.Address))
			if address.Number == nil {
				newAddress.WriteString("s/n, ")
			} else {
				newAddress.WriteString(fmt.Sprintf("%s,\n", *address.Number))
			}
			if address.Complement != nil {
				newAddress.WriteString(fmt.Sprintf("%s,\n", *address.Complement))
			}
			newAddress.WriteString(fmt.Sprintf("%s-%s\nCEP %s ", address.City, address.State, zipCode))
			return newAddress.String()
		}
	}
	return ""
}

func (person *Person) GetFormattedDocument() string {
	pattern := `^(\d{3})(\d{3})(\d{3})(\d{2})$`
	mask := "$1.$2.$3-$4"
	if len(person.Document) == 14 {
		pattern = `^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`
		mask = "$1.$2.$3/$4-$5"
	}
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(person.Document, mask)
}

func (person *Person) GetFistName() string {
	return strings.Split(person.Name, " ")[0]
}

func (PersonPhone) TableName() string {
	return "manager.people__phones"
}

func (PersonEmail) TableName() string {
	return "manager.people__emails"
}

func (PersonAddress) TableName() string {
	return "manager.people__addresses"
}

func (PersonExtra) TableName() string {
	return "manager.people__extras"
}

func (Person) TableName() string {
	return "manager.people"
}

func (phone *PersonPhone) BeforeCreate(tx *gorm.DB) (err error) {
	if phone.ID == uuid.Nil {
		phone.ID = uuid.New()
	}
	phone.CreatedAt = time.Now()
	return
}

func (email *PersonEmail) BeforeCreate(tx *gorm.DB) (err error) {
	if email.ID == uuid.Nil {
		email.ID = uuid.New()
	}
	email.CreatedAt = time.Now()
	return
}

func (address *PersonAddress) BeforeCreate(tx *gorm.DB) (err error) {
	if address.ID == uuid.Nil {
		address.ID = uuid.New()
	}
	address.CreatedAt = time.Now()
	return
}

func (extra *PersonExtra) BeforeCreate(tx *gorm.DB) (err error) {
	if extra.ID == uuid.Nil {
		extra.ID = uuid.New()
	}
	extra.CreatedAt = time.Now()
	return
}

func (person *Person) BeforeCreate(tx *gorm.DB) (err error) {
	person.ID = uuid.New()
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()
	return
}

func (phone *PersonPhone) BeforeUpdate(tx *gorm.DB) (err error) {
	phone.UpdatedAt = time.Now()
	return nil
}

func (email *PersonEmail) BeforeUpdate(tx *gorm.DB) (err error) {
	email.UpdatedAt = time.Now()
	return nil
}

func (address *PersonAddress) BeforeUpdate(tx *gorm.DB) (err error) {
	address.UpdatedAt = time.Now()
	return nil
}

func (extra *PersonExtra) BeforeUpdate(tx *gorm.DB) (err error) {
	extra.UpdatedAt = time.Now()
	return nil
}

func (person *Person) BeforeUpdate(tx *gorm.DB) (err error) {
	person.UpdatedAt = time.Now()
	return nil
}

type PersonRepository interface {
	Insert(person *Person) (insertedPerson *Person, err error)
	Read(id uuid.UUID) (person *Person, err error)
	ReadAll(page int, pageSize int) (people *[]Person, pagination *utils.Pagination, err error)
	ReadByEmail(email string) (person *Person, err error)
	ReadByPhoneNumber(countryCode int, areaCode int, phoneNumber int) (person *Person, err error)
	Update(person *Person) (updatedPersons *Person, err error)
	Delete(id uuid.UUID) (err error)
}
