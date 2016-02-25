package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*Pet Pet pet

swagger:model Pet
*/
type Pet struct {

	/* Category category
	 */
	Category *Category `json:"category,omitempty"`

	/* ID id
	 */
	ID *int64 `json:"id,omitempty"`

	/* Name name

	Required: true
	*/
	Name string `json:"name,omitempty"`

	/* PhotoUrls photo urls

	Required: true
	*/
	PhotoUrls []string `json:"photoUrls,omitempty" xml:"photoUrl"`

	/* pet status in the store
	 */
	Status *string `json:"status,omitempty"`

	/* Tags tags
	 */
	Tags []*Tag `json:"tags,omitempty" xml:"tag"`
}

// Validate validates this pet
func (m *Pet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhotoUrls(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pet) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *Pet) validatePhotoUrls(formats strfmt.Registry) error {

	if err := validate.Required("photoUrls", "body", m.PhotoUrls); err != nil {
		return err
	}

	for i := 0; i < len(m.PhotoUrls); i++ {

		if err := validate.RequiredString("photoUrls"+"."+strconv.Itoa(i), "body", string(m.PhotoUrls[i])); err != nil {
			return err
		}

	}

	return nil
}

var petStatusEnum []interface{}

func (m *Pet) validateStatusEnum(path, location string, value string) error {
	if petStatusEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["available","pending","sold"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			petStatusEnum = append(petStatusEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, petStatusEnum); err != nil {
		return err
	}
	return nil
}

func (m *Pet) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Pet) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if err := m.Tags[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
