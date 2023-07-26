// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserView -  The UserView object provides a user response object, as well as JSONPATHs to related objects provided by expanders.
type UserView struct {
	//  The User object provides all of the details for an user, as well as some configuration.
	//
	User *User `json:"user,omitempty"`
	//  JSONPATH expression indicating the location of the user objects of delegates of the current user in the expanded array.
	//
	DelegatedUserPath *string `json:"delegatedUserPath,omitempty"`
	//  JSONPATH expression indicating the location of directory objects in the expanded array.
	//
	DirectoriesPath *string `json:"directoriesPath,omitempty"`
	//  JSONPATH expression indicating the location of the user objects that managed the current user in the expanded array.
	//
	ManagersPath *string `json:"managersPath,omitempty"`
	//  JSONPATH expression indicating the location of the roles of the current user in the expanded array.
	//
	RolesPath *string `json:"rolesPath,omitempty"`
}

func (o *UserView) GetUser() *User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *UserView) GetDelegatedUserPath() *string {
	if o == nil {
		return nil
	}
	return o.DelegatedUserPath
}

func (o *UserView) GetDirectoriesPath() *string {
	if o == nil {
		return nil
	}
	return o.DirectoriesPath
}

func (o *UserView) GetManagersPath() *string {
	if o == nil {
		return nil
	}
	return o.ManagersPath
}

func (o *UserView) GetRolesPath() *string {
	if o == nil {
		return nil
	}
	return o.RolesPath
}
