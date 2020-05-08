# cck-sdk-go 介绍
此 SDK 用于首云的NAS云盘产品

# 安装

Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or project's Go module dependencies.

```go
go get github.com/capitalonline/cck-sdk-go
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.

```go
go get -u github.com/capitalonline/cck-sdk-go
```

# SDK 使用说明

容器启动 yaml 中配置环境变量（用户的认证信息）：

```reStructuredText
ACCESS_KEY_ID = "***"	
ACCESS_KEY_SECRET = "***"
```

代码调用示例：

```go
package main 

import (
    "github.com/capitalonline/cck-sdk-go/pkg/cck/nas"
    log "github.com/sirupsen/logrus"
) 

func main() {
    // MountNas usage example 
    nasID := "340d569c-7899-11ea-a06c-82b0d54620aa"
	clusterID := "3b02449e-8843-11ea-988c-0242ac11034c"
	res, err := nas.MountNas(nasID, clusterID)
    if err != nil {
		log.Errorf("Failed, err is: %s", err.Error())
	}
    task_id = res.data.TaskId
}
```

# SDK 参数说明

## 1.DescribeNasInstances 函数

**Action: **DescribeNasInstances
**描述：**查询文件存储(以下简称nas)详细信息
**请求方法：**POST
**请求参数：**

| **参数** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- |
| NasID | 否 | string | *** | nas实例id |
| NasName | 否 | string | *** | nas实例名字 |
| SiteID | 否 | string | *** | 节点id |
| ClusterID | 否 | string | *** | 集群id |
| UsageFlag | 否 | int | *** | 是否返回使用量 |
| PageNumber | 否 | int | *** | 页码 |
| PageSize | 否 | int | *** | 每页条目数 |
**返回数据：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |
| data | dict | {} | 返回数据字典 |
| total | interger | 2 | 返回列表条目数 |
| nas_info | list | [] | nas实例信息列表 |
| id | string | 5f19a292-759b-11ea-aaa6-0242ac110230 | nas实例id |
| name | string | test | nas实例名称 |
| site_id | string | c4089dcd-15c2-4caf-9d1e-874770a31880 | nas实例所属节点id |
| cluster_id | string | b62129e8-7596-11ea-b79c-0242ac11021e | nas实例挂载集群id |
| cluster_name | string | test | nas实例挂载集群名称 |
| disk_type | string | high_disk | nas实例磁盘类型 |
| iops | interger | 3000 | nas实例磁盘ipos |
| size | interger | 500 | nas实例磁盘大小 |
| status | string | ok | nas实例状态 |
| create_time | datetime | 2020-04-03T19:07:56.412827 | nas实例创建时间 |
| storage_vm_id | string | 12ab9419-f71e-4c66-bfeb-c88c3cd038ea | nas实例磁盘挂载虚机id |
| mount_point | string | 10.241.85.80 | nas实例磁盘挂载虚机ip |
| backup_disk_mount_path | string | /nfsshare | nas实例磁盘挂载路径 |
| usage | string | 0.08 | nas实例磁盘使用量 |
| usage_rate | string | 1% | nas实例磁盘使用量比率 |
**返回数据**

```json
{
  "code": "Success",
  "data": {
    "nas_info": [
      {
        "id": "5f19a292-759b-11ea-aaa6-0242ac110230",
        "name": "test-gz-hhh",
        "site_id": "c4089dcd-15c2-4caf-9d1e-874770a31880",
        "cluster_id": "b62129e8-7596-11ea-b79c-0242ac11021e",
        "cluster_name": "test-hys-gz",
        "disk_type": "high_disk",
        "iops": 3000,
        "size": 500,
        "status": "mounted",
        "create_time": "2020-04-03T19:07:56.412827",
        "storage_vm_id": "12ab9419-f71e-4c66-bfeb-c88c3cd038ea",
        "mount_point": "10.241.85.80",
        "backup_disk_mount_path": "/nfsshare",
        "status_str": "已挂载",
        "usage": "0.07",
        "usage_rate": "0.01%"
      },
      {
        "id": "c13dbc62-7576-11ea-88be-865e67cc729c",
        "name": "test001",
        "site_id": "3fd55550-8fd7-4634-84ba-1a9880ec1ce4",
        "cluster_id": "",
        "cluster_name": "",
        "disk_type": "high_disk",
        "iops": 3000,
        "size": 500,
        "status": "ok",
        "create_time": "2020-04-03T14:45:49.687300",
        "storage_vm_id": "ee20977c-d6ba-4ac4-ba88-b7d71d95fbbb",
        "mount_point": "",
        "backup_disk_mount_path": "/nfsshare",
        "status_str": "正常",
        "usage": "0.07",
        "usage_rate": "0.01%"
      }
    ],
    "total": 2
  }
}
```

## 2. CreateNas 函数

**Action: **CreateNas
**描述：**创建Nas实例
**请求方法：**POST
**请求参数：**

| **名称** | **外部名称** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- | --- |
| nas_site_id | SiteId | 是 | string | c4089dcd-15c2-4caf-9d1e-874770a31880 | 节点id |
| nas_name | NasName | 是 | string | test | nas存储名称 |
| disk_type | DiskTpye | 是 | string | high_disk | 磁盘类型 |
| disk_size | DiskSize | 是 | string | 500 | 磁盘大小 |

**返回参数：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |
| task_id | string | 17b2a6f8-78cd-11ea-b1f2-169433ed6e42 | 任务id |

**返回示例**
```json
{
  "code": "Success",
  "data": {
    "task_id": "17b2a6f8-78cd-11ea-b1f2-169433ed6e42"
  }
}
```

## 3. ResizeNas 函数

**Action: **ResizeNas
**描述：**扩容Nas实例大小
**请求方法：**POST
**请求参数：**

| **名称** | **外部名称** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- | --- |
| nas_site_id | SiteId | 是 | string | c4089dcd-15c2-4caf-9d1e-874770a31880 | 节点id |
| nas_id | NasId | 是 | string | 5f19a292-759b-11ea-aaa6-0242ac110230 | Nas实例id |
| disk_size | DiskSize | 是 | string | 1000 | Nas实例磁盘大小（单位为G） |

**返回参数：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |
| task_id | string | 17b2a6f8-78cd-11ea-b1f2-169433ed6e42 | 任务id |

**返回示例**
```json
{
  "code": "Success",
  "data": {
    "task_id": "17b2a6f8-78cd-11ea-b1f2-169433ed6e42"
  }
}
```

## 4. DeleteNas 函数

**Action: **DeleteNas
**描述：**删除Nas实例
**请求方法：**POST
**请求参数：**

| **名称** | **外部名称** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- | --- |
| nas_id | NasId | 是 | string | 5f19a292-759b-11ea-aaa6-0242ac110230 | Nas实例id |

**返回参数：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |

**返回示例**
```json
{
  "code": "Success",
}
```

## 5. MountNas 函数

**Action: **MountNas
**描述：**挂载Nas实例
**请求方法：**POST
**请求参数：**

| **名称** | **外部名称** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- | --- |
| nas_id | NasId | 是 | string | 5f19a292-759b-11ea-aaa6-0242ac110230 | Nas实例id |
| cluster_id | ClusterId | 是 | string | b62129e8-7596-11ea-b79c-0242ac11021e | 集群id |

**返回参数：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |
| task_id | stirng | 17b2a6f8-78cd-11ea-b1f2-169433ed6e42 | 任务id |

**返回示例**
```json
{
  "code": "Success",
  "data": {
    "task_id": "17b2a6f8-78cd-11ea-b1f2-169433ed6e42"
  }
}
```

## 6. UmountNas 函数

**Action: **UmountNas
**描述：**卸载Nas实例
**请求方法：**POST
**请求参数：**

| **名称** | **外部名称** | **是否必选** | **类型** | **实例值** | **描述** |
| --- | --- | --- | --- | --- | --- |
| nas_id | NasId | 是 | string | 5f19a292-759b-11ea-aaa6-0242ac110230 | Nas实例id |

**返回参数：**

| **名称** | **类型** | **示例值** | **描述** |
| --- | --- | --- | --- |
| code | string | Success | 错误码 |
| task_id | stirng | 17b2a6f8-78cd-11ea-b1f2-169433ed6e42 | 任务id |

**返回示例**
```json
{
  "code": "Success",
  "data": {
    "task_id": "17b2a6f8-78cd-11ea-b1f2-169433ed6e42"
  }
}
```

# License

This SDK is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see LICENSE for more information.

