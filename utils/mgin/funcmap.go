/*
 * Copyright (c) 2017 - 2020, 奶爸<hi@nai.ba>
 * All rights reserved.
 */

package mgin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"runtime"
	"strings"
	"time"

	"github.com/naiba/nocd"
)

//FuncMap 自定义模板函数
func FuncMap(pipelineService nocd.PipelineService, pipelogService nocd.PipeLogService, webhookService nocd.WebhookService) template.FuncMap {
	return template.FuncMap{
		"RepoPipelines": func(rid uint) []nocd.Pipeline {
			return pipelineService.RepoPipelines(&nocd.Repository{ID: rid})
		},
		"UserPipelines": func(uid uint) []nocd.Pipeline {
			var u nocd.User
			u.ID = uid
			return pipelineService.UserPipelines(&u)
		},
		"PipelineWebhooks": func(pid uint) []nocd.Webhook {
			return webhookService.PipelineWebhooks(&nocd.Pipeline{ID: pid})
		},
		"LastServerLog": func(rid uint) nocd.PipeLog {
			return pipelogService.LastServerLog(rid)
		},
		"JSON": func(obj interface{}) string {
			b, _ := json.Marshal(obj)
			return string(b)
		},
		"LastPipelineLog": func(pid uint) nocd.PipeLog {
			return pipelogService.LastPipelineLog(pid)
		},
		"TimeDiff": func(t1, t2 time.Time) string {
			if t2.IsZero() {
				return "正在运行"
			}
			sec := t2.Sub(t1).Seconds()
			if sec < 60 {
				return fmt.Sprintf(" %.0f 秒", sec)
			}
			if sec < 60*60 {
				return fmt.Sprintf(" %.0f 分钟", sec/60)
			}
			if sec < 60*60*24 {
				return fmt.Sprintf(" %.0f 小时", sec/60/60)
			}
			if sec < 60*60*24*30 {
				return fmt.Sprintf(" %.0f 天", sec/60/60/24)
			}
			if sec < 60*60*24*30*12 {
				return fmt.Sprintf(" %.0f 个月", sec/60/60/24/30)
			}
			return fmt.Sprintf(" %.0f 年", sec/60/60/24/30/12)
		},
		"Now": func() time.Time {
			return time.Now().In(nocd.Loc)
		},
		"TimeFormat": func(t time.Time) string {
			return t.In(nocd.Loc).Format("2006-01-02 15:04:05")
		},
		"HasPrefix": strings.HasPrefix,
		"MathSub": func(o, n int64) int64 {
			return o - n
		},
		"MathAdd": func(o, n int64) int64 {
			return o + n
		},
		"NumGoroutine": runtime.NumGoroutine,
	}
}
