// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/stark-sim/cas/pkg/ent/role"
	"github.com/stark-sim/cas/pkg/ent/schema"
	"github.com/stark-sim/cas/pkg/ent/user"
	"github.com/stark-sim/cas/pkg/ent/userrole"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedBy is the schema descriptor for created_by field.
	roleDescCreatedBy := roleMixinFields0[1].Descriptor()
	// role.DefaultCreatedBy holds the default value on creation for the created_by field.
	role.DefaultCreatedBy = roleDescCreatedBy.Default.(int64)
	// roleDescUpdatedBy is the schema descriptor for updated_by field.
	roleDescUpdatedBy := roleMixinFields0[2].Descriptor()
	// role.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	role.DefaultUpdatedBy = roleDescUpdatedBy.Default.(int64)
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[3].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[4].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescDeletedAt is the schema descriptor for deleted_at field.
	roleDescDeletedAt := roleMixinFields0[5].Descriptor()
	// role.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	role.DefaultDeletedAt = roleDescDeletedAt.Default.(time.Time)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[0].Descriptor()
	// role.DefaultName holds the default value on creation for the name field.
	role.DefaultName = roleDescName.Default.(string)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleMixinFields0[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() int64)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(int64)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(int64)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[5].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
	userroleMixin := schema.UserRole{}.Mixin()
	userroleMixinFields0 := userroleMixin[0].Fields()
	_ = userroleMixinFields0
	userroleFields := schema.UserRole{}.Fields()
	_ = userroleFields
	// userroleDescCreatedBy is the schema descriptor for created_by field.
	userroleDescCreatedBy := userroleMixinFields0[1].Descriptor()
	// userrole.DefaultCreatedBy holds the default value on creation for the created_by field.
	userrole.DefaultCreatedBy = userroleDescCreatedBy.Default.(int64)
	// userroleDescUpdatedBy is the schema descriptor for updated_by field.
	userroleDescUpdatedBy := userroleMixinFields0[2].Descriptor()
	// userrole.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	userrole.DefaultUpdatedBy = userroleDescUpdatedBy.Default.(int64)
	// userroleDescCreatedAt is the schema descriptor for created_at field.
	userroleDescCreatedAt := userroleMixinFields0[3].Descriptor()
	// userrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	userrole.DefaultCreatedAt = userroleDescCreatedAt.Default.(func() time.Time)
	// userroleDescUpdatedAt is the schema descriptor for updated_at field.
	userroleDescUpdatedAt := userroleMixinFields0[4].Descriptor()
	// userrole.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userrole.DefaultUpdatedAt = userroleDescUpdatedAt.Default.(func() time.Time)
	// userrole.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userrole.UpdateDefaultUpdatedAt = userroleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userroleDescDeletedAt is the schema descriptor for deleted_at field.
	userroleDescDeletedAt := userroleMixinFields0[5].Descriptor()
	// userrole.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	userrole.DefaultDeletedAt = userroleDescDeletedAt.Default.(time.Time)
	// userroleDescID is the schema descriptor for id field.
	userroleDescID := userroleMixinFields0[0].Descriptor()
	// userrole.DefaultID holds the default value on creation for the id field.
	userrole.DefaultID = userroleDescID.Default.(func() int64)
}
