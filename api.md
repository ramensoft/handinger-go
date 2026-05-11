# Workers

Params Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#CreateWorkerParam">CreateWorkerParam</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#UpdateWorkerParam">UpdateWorkerParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#DeleteWorkerResponse">DeleteWorkerResponse</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Worker">Worker</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerTemplate">WorkerTemplate</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerGetEmailResponse">WorkerGetEmailResponse</a>

Methods:

- <code title="post /api/workers">client.Workers.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerNewParams">WorkerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerTemplate">WorkerTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workers/{workerId}">client.Workers.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerGetParams">WorkerGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Worker">Worker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /api/workers/{workerId}">client.Workers.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerUpdateParams">WorkerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerTemplate">WorkerTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workers/{workerId}">client.Workers.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#DeleteWorkerResponse">DeleteWorkerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workers/{workerId}/email">client.Workers.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerService.GetEmail">GetEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerGetEmailResponse">WorkerGetEmailResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Schedules

Response Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleUnion">WorkerScheduleUnion</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleListResponse">WorkerScheduleListResponse</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleCancelResponse">WorkerScheduleCancelResponse</a>

Methods:

- <code title="post /api/workers/{workerId}/schedules">client.Workers.Schedules.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleNewParams">WorkerScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleUnion">WorkerScheduleUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workers/{workerId}/schedules">client.Workers.Schedules.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleListResponse">WorkerScheduleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workers/{workerId}/schedules/{scheduleId}">client.Workers.Schedules.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scheduleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleCancelParams">WorkerScheduleCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerScheduleCancelResponse">WorkerScheduleCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Webhooks

Params Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#UpdateWebhookParam">UpdateWebhookParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Webhook">Webhook</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WebhookExecution">WebhookExecution</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WebhookExecutionList">WebhookExecutionList</a>

Methods:

- <code title="get /api/workers/{workerId}/webhook">client.Workers.Webhooks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /api/workers/{workerId}/webhook">client.Workers.Webhooks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookUpdateParams">WorkerWebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/workers/{workerId}/webhook">client.Workers.Webhooks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/workers/{workerId}/webhook/executions">client.Workers.Webhooks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookService.ListExecutions">ListExecutions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookListExecutionsParams">WorkerWebhookListExecutionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WebhookExecutionList">WebhookExecutionList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/workers/{workerId}/webhook/regenerate-token">client.Workers.Webhooks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#WorkerWebhookService.RegenerateToken">RegenerateToken</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tasks

Params Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#CreateTaskParam">CreateTaskParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#DeleteTaskResponse">DeleteTaskResponse</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Task">Task</a>
- <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskWithTurns">TaskWithTurns</a>

Methods:

- <code title="post /api/tasks">client.Tasks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskNewParams">TaskNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#Worker">Worker</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/tasks/{taskId}">client.Tasks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskWithTurns">TaskWithTurns</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/tasks/{taskId}">client.Tasks.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#TaskService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, taskID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go">handinger</a>.<a href="https://pkg.go.dev/github.com/Ramensoft/handinger-go#DeleteTaskResponse">DeleteTaskResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
