package distribution

import (
	"../errors"
)

type Distributor struct {
	Name        string
	permissions PermissionMatrix
	parent      *Distributor
}

// initialize the distributor
func (distributor *Distributor) Initialize(name string, parent *Distributor) errors.ApplicationError {
	distributor.Name = name
	distributor.parent = parent
	distributor.permissions = loadBasePermissions()
	return nil
}

// include the location to the distributor permissions
func (distributor *Distributor) Include(location string) errors.ApplicationError {
	// if it is NOT a sub-distributor, then it has scope to all locations
	if distributor.parent == nil {
		return distributor.permissions.Include(location)
	} else { // otherwise, the scope is limited to scope of the parent distributor
		if distributor.parent.permissions.IsAllowed(location) {
			return distributor.permissions.Copy(location, distributor.parent.permissions)
		} else {
			return DistributionScopeError(location)
		}
	}
}

// exclude the location to the distributor permissions
func (distributor *Distributor) Exclude(location string) errors.ApplicationError {
	// no need to check the scope for exclude
	return distributor.permissions.Exclude(location)
}

// query if the distributor can distribute in a location
func (distributor *Distributor) CanDistribute(location string) bool {
	// query the location in permission matrix
	return distributor.permissions.IsAllowed(location)
}

// TODO(ilayaraja): Do not expose
// func (distributor *Distributor) Permissions() PermissionMatrix {
// 	return distributor.permissions
// }
