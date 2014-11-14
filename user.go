// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// User represents a API user.
type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}
