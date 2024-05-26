# Manager Service

Service for manage business workflows.

### Environment variables
|**Variable**|**description**|
|--|--|
|MODE|application running mode (default production)|
|NAME|service display name (default manager)|
|PORT|service port (default 8080)|
|--|--|
|DATABASE_HOST|database host path|
|DATABASE_PORT|database port (default 5432)|
|DATABASE_NAME|database name|
|DATABASE_USERNAME|database username|
|DATABASE_PASSWORD|database password|
|--|--|
|AUTH_HOST|keycloak host url|
|AUTH_REALM|keycloak realm|
|AUTH_CLIENT|keycloak client id|
|AUTH_CLIENT_SECRET|keycloak client secret|
|AUTH_PUBLIC_KEY|keycloak rsa public key|
|--|--|

## Technologies
- Go (Golang)
- PostgreSql

### Software architecture
##### Clean architecture based:
- attachment - `attachment file management`
- board - `project board management`
- content - `item content management`
- item - `step item management`
- label - ``label management`
- project - ``project management`
- step - ``board step management`
- task - ``task management`

