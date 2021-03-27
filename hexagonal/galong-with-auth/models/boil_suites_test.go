// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Accounts", testAccounts)
	t.Run("Permissions", testPermissions)
	t.Run("RoleUsers", testRoleUsers)
	t.Run("Roles", testRoles)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Permissions", testPermissionsDelete)
	t.Run("RoleUsers", testRoleUsersDelete)
	t.Run("Roles", testRolesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Permissions", testPermissionsQueryDeleteAll)
	t.Run("RoleUsers", testRoleUsersQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Permissions", testPermissionsSliceDeleteAll)
	t.Run("RoleUsers", testRoleUsersSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Permissions", testPermissionsExists)
	t.Run("RoleUsers", testRoleUsersExists)
	t.Run("Roles", testRolesExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Permissions", testPermissionsFind)
	t.Run("RoleUsers", testRoleUsersFind)
	t.Run("Roles", testRolesFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Permissions", testPermissionsBind)
	t.Run("RoleUsers", testRoleUsersBind)
	t.Run("Roles", testRolesBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Permissions", testPermissionsOne)
	t.Run("RoleUsers", testRoleUsersOne)
	t.Run("Roles", testRolesOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Permissions", testPermissionsAll)
	t.Run("RoleUsers", testRoleUsersAll)
	t.Run("Roles", testRolesAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Permissions", testPermissionsCount)
	t.Run("RoleUsers", testRoleUsersCount)
	t.Run("Roles", testRolesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accounts", testAccountsHooks)
	t.Run("Permissions", testPermissionsHooks)
	t.Run("RoleUsers", testRoleUsersHooks)
	t.Run("Roles", testRolesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Permissions", testPermissionsInsert)
	t.Run("Permissions", testPermissionsInsertWhitelist)
	t.Run("RoleUsers", testRoleUsersInsert)
	t.Run("RoleUsers", testRoleUsersInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PermissionToRoleUsingRole", testPermissionToOneRoleUsingRole)
	t.Run("RoleUserToRoleUsingRole", testRoleUserToOneRoleUsingRole)
	t.Run("RoleUserToAccountUsingUser", testRoleUserToOneAccountUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AccountToUserRoleUsers", testAccountToManyUserRoleUsers)
	t.Run("RoleToPermissions", testRoleToManyPermissions)
	t.Run("RoleToRoleUsers", testRoleToManyRoleUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PermissionToRoleUsingPermissions", testPermissionToOneSetOpRoleUsingRole)
	t.Run("RoleUserToRoleUsingRoleUsers", testRoleUserToOneSetOpRoleUsingRole)
	t.Run("RoleUserToAccountUsingUserRoleUsers", testRoleUserToOneSetOpAccountUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("PermissionToRoleUsingPermissions", testPermissionToOneRemoveOpRoleUsingRole)
	t.Run("RoleUserToRoleUsingRoleUsers", testRoleUserToOneRemoveOpRoleUsingRole)
	t.Run("RoleUserToAccountUsingUserRoleUsers", testRoleUserToOneRemoveOpAccountUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AccountToUserRoleUsers", testAccountToManyAddOpUserRoleUsers)
	t.Run("RoleToPermissions", testRoleToManyAddOpPermissions)
	t.Run("RoleToRoleUsers", testRoleToManyAddOpRoleUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AccountToUserRoleUsers", testAccountToManySetOpUserRoleUsers)
	t.Run("RoleToPermissions", testRoleToManySetOpPermissions)
	t.Run("RoleToRoleUsers", testRoleToManySetOpRoleUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AccountToUserRoleUsers", testAccountToManyRemoveOpUserRoleUsers)
	t.Run("RoleToPermissions", testRoleToManyRemoveOpPermissions)
	t.Run("RoleToRoleUsers", testRoleToManyRemoveOpRoleUsers)
}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Permissions", testPermissionsReload)
	t.Run("RoleUsers", testRoleUsersReload)
	t.Run("Roles", testRolesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Permissions", testPermissionsReloadAll)
	t.Run("RoleUsers", testRoleUsersReloadAll)
	t.Run("Roles", testRolesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Permissions", testPermissionsSelect)
	t.Run("RoleUsers", testRoleUsersSelect)
	t.Run("Roles", testRolesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Permissions", testPermissionsUpdate)
	t.Run("RoleUsers", testRoleUsersUpdate)
	t.Run("Roles", testRolesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Permissions", testPermissionsSliceUpdateAll)
	t.Run("RoleUsers", testRoleUsersSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
}