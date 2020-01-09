/*
Copyright 2020 Cloudera, Inc.  All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package security

import (
	"os/user"
	"time"
)

// Get the cache and use that to resolve all user requests
func GetUserGroupCacheOS() *UserGroupCache {
	return &UserGroupCache{
		ugs:           map[string]*UserGroup{},
		interval:      cleanerInterval * time.Second,
		lookup:        user.Lookup,
		lookupGroupID: user.LookupGroupId,
		groupIds:      wrappedGroupIds,
	}
}

// wrapper function to allow easy testing of the cache
func wrappedGroupIds(osUser *user.User) ([]string, error) {
	return osUser.GroupIds()
}
