/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"configcenter/src/auth/extensions"
	"configcenter/src/common/blog"
	"configcenter/src/common/condition"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/topo_server/core/types"
	"fmt"
)

// CreateObjectAttribute create a new object attribute
func (s *Service) CreateObjectAttribute(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	attr, err := s.Core.AttributeOperation().CreateObjectAttribute(params, data)
	if nil != err {
		return nil, err
	}

	// auth: register resource
	attribute := attr.Attribute()
	authManager := extensions.NewAuthManager(s.Engine.CoreAPI, s.Auth.Authorizer, params.Err)
	if err := authManager.RegisterModelAttribute(params.Context, params.Header, *attribute); err != nil {
		return nil, fmt.Errorf("register model attribute to auth failed, err: %+v", err)
	}
	
	return attr.ToMapStr()
}

// SearchObjectAttribute search the object attributes
func (s *Service) SearchObjectAttribute(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	cond := condition.CreateCondition()
	data.Remove(metadata.PageName)
	if err := cond.Parse(data); nil != err {
		blog.Errorf("search object attribute, but failed to parse the data into condition, err: %v", err)
		return nil, err
	}
	cond.Field(metadata.AttributeFieldIsSystem).NotEq(true)
	cond.Field(metadata.AttributeFieldIsAPI).NotEq(true)
	return s.Core.AttributeOperation().FindObjectAttributeWithDetail(params, cond)
}

// UpdateObjectAttribute update the object attribute
func (s *Service) UpdateObjectAttribute(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {
	paramPath := mapstr.MapStr{}
	paramPath.Set("id", pathParams("id"))
	id, err := paramPath.Int64("id")
	if nil != err {
		blog.Errorf("[api-att] failed to parse the path params id(%s), error info is %s ", pathParams("id"), err.Error())
		return nil, err
	}

	data.Remove(metadata.BKMetadata)

	err = s.Core.AttributeOperation().UpdateObjectAttribute(params, data, id)

	// auth: update registered resource
	authManager := extensions.NewAuthManager(s.Engine.CoreAPI, s.Auth.Authorizer, params.Err)
	if err := authManager.UpdateRegisteredModelAttributeByID(params.Context, params.Header, id); err != nil {
		return nil, fmt.Errorf("update registered model attribute to auth failed, err: %+v", err)
	}

	return nil, err
}

// DeleteObjectAttribute delete the object attribute
func (s *Service) DeleteObjectAttribute(params types.ContextParams, pathParams, queryParams ParamsGetter, data mapstr.MapStr) (interface{}, error) {

	paramPath := mapstr.MapStr{}
	paramPath.Set("id", pathParams("id"))
	id, err := paramPath.Int64("id")
	if nil != err {
		blog.Errorf("[api-att] failed to parse the path params id(%s), error info is %s ", pathParams("id"), err.Error())
		return nil, err
	}

	cond := condition.CreateCondition()
	cond.Field(metadata.AttributeFieldSupplierAccount).Eq(params.SupplierAccount)
	cond.Field(metadata.AttributeFieldID).Eq(id)

	data.Remove(metadata.BKMetadata)

	err = s.Core.AttributeOperation().DeleteObjectAttribute(params, cond)

	// auth: update registered resource
	authManager := extensions.NewAuthManager(s.Engine.CoreAPI, s.Auth.Authorizer, params.Err)
	if err := authManager.DeregisterModelAttributeByID(params.Context, params.Header, id); err != nil {
		return nil, fmt.Errorf("update registered model attribute to auth failed, err: %+v", err)
	}

	return nil, err
}
