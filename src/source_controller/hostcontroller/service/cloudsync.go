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
	"configcenter/src/common"
	"configcenter/src/common/blog"
	meta "configcenter/src/common/metadata"
	"configcenter/src/common/util"
	"context"
	"encoding/json"
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) AddCloud(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ownerID := util.GetOwnerID(pheader)
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	input := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&input); err != nil {
		blog.Errorf("add cloud sync task failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	input[common.CreateTimeField] = time.Now()
	input = util.SetModOwner(input, ownerID)

	err := s.Logics.CreateCloudTask(ctx, input)
	if err != nil {
		blog.Errorf("create cloud sync data: %v error: %v", input, err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCloudSyncCreateFail)})
		return
	}

	result := make(map[string]interface{})
	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     result,
	})
}

func (s *Service) TaskNameCheck(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ownerID := util.GetOwnerID(pheader)
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	input := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&input); err != nil {
		blog.Errorf("task name check failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	condition := common.KvMap{"bk_task_name": input["bk_task_name"]}
	condition = util.SetModOwner(condition, ownerID)
	num, err := s.Instance.Table(common.BKTableNameCloudTask).Find(condition).Count(ctx)
	if err != nil {
		blog.Error("get task name [%s] failed, err: %v", input["bk_task_name"], err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	//blog.Info("task name uniqueness check")
	//blog.Debug("num = %v", num)
	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     num,
	})
}

func (s *Service) DeleteCloudTask(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	taskID := req.PathParameter("taskID")
	intTaskID, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		blog.Errorf("string to int64 failed with err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommParamsIsInvalid)})
		return
	}

	params := common.KvMap{"bk_task_id": intTaskID}
	blog.Debug("params: %v", params)

	if err := s.Instance.Table(common.BKTableNameCloudTask).Delete(ctx, params); err != nil {
		blog.Errorf("delete failed err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBDeleteFailed)})
		return
	}

	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     "success",
	})
}

func (s *Service) SearchCloudTask(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	condition := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&condition); err != nil {
		blog.Errorf("add cloud sync task failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	result := make([]map[string]interface{}, 0)
	err := s.Instance.Table(common.BKTableNameCloudTask).Find(condition).All(ctx, &result)
	if err != nil {
		blog.Error("get failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	num, err := s.Instance.Table(common.BKTableNameCloudTask).Find(condition).Count(ctx)
	if err != nil {
		blog.Error("get task name [%s] failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	//blog.Debug("num: %v", num)
	//blog.Debug("result: %v", result)
	resp.WriteEntity(meta.FavoriteResult{
		Count: num,
		Info:  result,
	})

}

func (s *Service) UpdateCloudTask(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), pheader)

	data := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&data); err != nil {
		blog.Errorf("update cloud task failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	params := common.KvMap{"bk_task_id": data["bk_task_id"]}
	//num, err := s.Instance.Table(common.BKTableNameCloudTask).Find(params).Count(ctx)

	err := s.Instance.Table(common.BKTableNameCloudTask).Update(ctx, params, data)
	if nil != err {
		blog.Error("update cloud task failed, error information is %s, params:%v", err.Error(), params)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBUpdateFailed)})
		return
	}

	resp.WriteEntity(meta.NewSuccessResp(nil))
}

func (s *Service) ResourceConfirm(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ownerID := util.GetOwnerID(pheader)
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	input := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&input); err != nil {
		blog.Errorf("add cloud sync task failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	input[common.CreateTimeField] = time.Now()
	input = util.SetModOwner(input, ownerID)

	err := s.Logics.CreateResourceConfirm(ctx, input)
	if err != nil {
		blog.Errorf("create cloud sync data: %v error: %v", input, err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCloudSyncCreateFail)})
		return
	}

	blog.Debug("input222: %v", input)

	result := make(map[string]interface{})
	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     result,
	})

}

func (s *Service) SearchConfirm(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	condition := make(map[string]interface{})
	result := make([]map[string]interface{}, 0)
	err := s.Instance.Table(common.BKTableNameCloudResourceConfirm).Find(condition).All(ctx, &result)
	if err != nil {
		blog.Error("get failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	num, errN := s.Instance.Table(common.BKTableNameCloudResourceConfirm).Find(condition).Count(ctx)
	if errN != nil {
		blog.Error("get task name [%s] failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	//blog.Debug("result: %v", result)
	resp.WriteEntity(meta.FavoriteResult{
		Count: num,
		Info:  result,
	})
}

func (s *Service) DeleteConfirm(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	resourceID := req.PathParameter("resourceID")
	intResourceID, err := strconv.ParseInt(resourceID, 10, 64)
	if err != nil {
		blog.Errorf("string to int64 failed with err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommParamsIsInvalid)})
		return
	}

	params := common.KvMap{"bk_resource_id": intResourceID}
	if err := s.Instance.Table(common.BKTableNameCloudResourceConfirm).Delete(ctx, params); err != nil {
		blog.Errorf("delete failed err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBDeleteFailed)})
		return
	}

	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     "success",
	})
}

func (s *Service) CloudHistory(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	input := make(map[string]interface{})
	if err := json.NewDecoder(req.Request.Body).Decode(&input); err != nil {
		blog.Errorf("add cloud sync task failed with decode body err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommJSONUnmarshalFailed)})
		return
	}

	err := s.Logics.CreateCloudHistory(ctx, input)
	if err != nil {
		blog.Errorf("create cloud history data: %v error: %v", input, err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCloudHistoryCreateFail)})
		return
	}

	result := make(map[string]interface{})
	resp.WriteEntity(meta.Response{
		BaseResp: meta.SuccessBaseResp,
		Data:     result,
	})
}

func (s *Service) SearchSyncHistory(req *restful.Request, resp *restful.Response) {
	pheader := req.Request.Header
	defErr := s.Core.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(pheader))
	ctx := util.GetDBContext(context.Background(), req.Request.Header)

	taskID := req.PathParameter("taskID")
	blog.Debug("taskID: %v", taskID)
	intResourceID, errInt := strconv.ParseInt(taskID, 10, 64)
	if errInt != nil {
		blog.Errorf("string to int64 failed.")
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommParamsIsInvalid)})
		return
	}

	condition := make(map[string]interface{})
	condition["bk_task_id"] = intResourceID

	result := make([]map[string]interface{}, 0)
	err := s.Instance.Table(common.BKTableNameCloudHistory).Find(condition).All(ctx, &result)
	if err != nil {
		blog.Error("get failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	num, errN := s.Instance.Table(common.BKTableNameCloudHistory).Find(condition).Count(ctx)
	if errN != nil {
		blog.Error("get task name [%s] failed, err: %v", err)
		resp.WriteError(http.StatusBadRequest, &meta.RespError{Msg: defErr.Error(common.CCErrCommDBSelectFailed)})
		return
	}

	resp.WriteEntity(meta.FavoriteResult{
		Count: num,
		Info:  result,
	})
}
