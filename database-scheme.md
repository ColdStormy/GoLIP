# Database tables

|*users*| |
| --- | --- |
|`id`       | primary key |
|`display_name`         | string |
|`name`         | string |
|`surname`      | string |
|`email` | string |


|*sessions*| _in-memory db_ |
| -- | -- |
|`user_id` | foreign key to *users* |
|`id` | primary key |
|`expiration_date` | date |
|`access_key` | string |
|`refresh_key` | string |


|*roles* | |
| -- | -- |
|`id` | primary key |
|`name` | string |


|*user_roles* | |
| -- | -- |
|`user_id` | foreign key to *users* |
|`role_id` | foreign key to *roles* |

