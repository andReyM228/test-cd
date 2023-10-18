# POSTGRESQL

- /app/{filepath}
- change_language
- /{url}
- login
- signup
- update_profile
- get_author_diamon_name
- get_user_info (get_author_info, get_info)
- get_diamon_info
- get_qr_code (get_cards_qr_code, get_qr_code)
- get_companies_list
- create_company
- get_company_info

---

# MONGODB

// for chat

### POST get_cards = get_levels_by_mission_id

#### request:

```json
{
  "mission_id": "qwerty";
  "user_id": "qwerty";
  "lobby_id": "qwerty";
}
```

#### response:

```json
{
  "levels": [
    {
      "id": "qwerty",
      "title": "qwerty",
      "description": "qwerty",
      "case_content": "qwerty",
      "level_progress"{
        "completed": true,
        "num_diamonds": 0,
      }
    }
  ]
}
```

### POST get_history

#### request:

```json
{
  "lobby_id": "qwerty",
  "user_id": "qwerty",
  "level_id": "qwerty",
  "mission_id": "qwerty"
}
```

#### response:

```json
{
  "history": [
    {
      "role": "string",
      "content": "string"
    }
  ]
}
```

### POST reorder_mission_question = reorder_levels

#### request:

```json
{
  "mission_id": "qwerty",
  "levels": ["id1", "id2"]
}
```

#### response:

```json
{
  "levels": [
    {
      "id": "qwerty",
      "title": "qwerty",
      "description": "qwerty",
      "case_content": "qwerty",
      "next_level_info": "qwerty",
      "type": "qwerty"
    }
  ]
}
```

### POST get_mission_content = get_mission_by_id

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "id": "string",
  "author_id": "string",
  "company_id": "string",
  "title": "string",
  "methodology_id": "string",
  "image_url": "string",
  "intro": [
    {
      "id": "string",
      "title": "string",
      "content": "string",
      "type": "string",
      "links": ["string"],
      "materials": [
        {
          "filename": "string",
          "file_url": "string"
        }
      ]
    }
  ],
  "materials": [
    {
      "id": "string",
      "filename": "string",
      "vector_count": 0
    }
  ],
  "levels": [
    {
      "id": "string",
      "title": "string",
      "description": "string",
      "case_content": "string",
      "type": "string"
    }
  ]
}
```

### POST delete_level (delete_level, delete_question, delete_case) = delete_level_by_id

#### request:

```json
{
  "mission_id": "qwerty",
  "level_id": "qwerty"
}
```

#### response:

```json
{}
```

### POST add_level (add_new_level_with_case, add_mission_question)

#### request:

```json
{
  "index": "0",
  "mission_id": "qwerty",
  "level": {
    "title": "string",
    "description": "string",
    "case_content": "string",
    "type": "string"
  }
}
```

#### response:

```json
{
  "id": "string",
  "title": "string",
  "description": "string",
  "case_content": "string",
  "type": "string"
}
```

### POST update_case (change_case) = update_level

#### request:

```json
{
  "mission_id": "qwerty",
  "level": {
    "id": "qwerty",
    "title": "qwerty",
    "description": "qwerty",
    "case_content": "qwerty",
  }
}
```

#### response:

```json
{
  "id": "qwerty",
  "title": "qwerty",
  "description": "qwerty",
  "case_content": "qwerty",
  "next_level_info": "qwerty",
  "type": "qwerty"
}
```

### POST create_lobby: takes courses_id = add_lobby

#### request:

```json
{
  "author_id": "qwerty",
  "title": "qwerty",
  "course_id": "qwerty",
  "mission_id": "qwerty"
}
```

#### response:

```json
{
  "id": "qwerty"
}
```

### POST get_lobby_by_id

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "id": "qwerty",
  "author_id": "qwerty",
  "name": "qwerty",
  "title": "qwerty",
  "course_id": "qwerty",
  "missions": {
    "mission_id": {
      "summary": "qwerty",
      "users": [
        {
          "user_id": "qwerty",
          "summary": "qwerty",
          "total_diamonds": 3,
          "image_url": "qwerty",
          "last_completed_level": 3
        }
      ]
    }
  }
  // "missions": [{
  //   "id": "qwerty",
  //   "summary": "qwerty",
  //   "users": [
  //     {
  //       "user_id": "qwerty",
  //       "summary": "qwerty",
  //       "total_diamonds": 3,
  //       "image_url": "qwerty",
  //       "last_completed_level": 3
  //     }
  //   ]
  // }]
}
```

### POST get_lobbies_by_author_id = get_lobbies_by_user_id (для игрока)

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "lobbies": [
    {
      "id": "string",
      "author_id": "string",
      "name": "string",
      "course_id": "string",
      "missions": [{
        "id": "string",
        "user_id": "string",
        "summary": "string",
        "completed": true,
      }]
    }
  ]
}
```

### POST get_lobbies_by_author_id = get_lobbies_by_mission_id

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "lobbies": [
    {
      "id": "string",
      "author_id": "string",
      "name": "string",
      "course_id": "string",
      "missions": [{
          "id": "string",
          "summary": "string",
          "users": [
            {
              "user_id": "string",
              "summary": "string",
              "total_diamonds": 0,
              "image_url": "string",
              "last_completed_level": 0
            }
          ]
      }]
    }
  ]
}
```

### POST delete_lobby

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{}
```

### POST get_missions_by_author_id

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "missions": [
    {
      "id": "string",
      "author_id": "string",
      "title": "string",
      "methodology_id": "string",
      "image_url": "string",
    }
  ]
}
```

### POST get_courses_by_author_id

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "courses": [
    {
      "id": "string",
      "title": "string",
      "author_id": "string",
      "company_id": "string",
      "created_at": "string",
      "missions_ids": ["string"]
    }
  ]
}
```

### POST complete_level

#### request:

```json
{
  "lobby_id": "qwerty",
  "user_id": "qwerty",
  "level_id": "qwerty",
  "mission_id": "qwerty"
}
```

#### response:

```json
{}
```

### POST change_intro_content

#### request:

```json
{
  "mission_id": "qwerty",
  "intro_id": "qwerty",
  "content": "qwerty"
}
```

#### response:

```json
{}
```

### POST add_intro_links

#### request:

```json
{
  "mission_id": "qwerty",
  "intro_id": "qwerty",
  "links": ["qwerty"]
}
```

#### response:

```json
{}
```

### POST delete_intro_link

#### request:

```json
{
  "mission_id": "qwerty",
  "intro_id": "qwerty",
  "link_index": "qwerty"
}

#### response:

```json
{}
```

### POST add_intro_materials

#### request:

```json
{
  "mission_id": "qwerty",
  "intro_id": "qwerty",
  "materials": [
    {
      "filename": "qwerty",
      "file_url": "qwerty"
    }
  ]
}
```

#### response:

```json
{}
```

### POST delete_intro_material

#### request:

```json
{
    "mission_id": "qwerty",
    "intro_id": "qwerty",
    "material_id": "qwerty"
}
```

#### response:

```json
{}
```

### get_users_by_lobby (wss) implemented

---

# BOTH

- start

---

# NONE

- change_profile_image

### lobby_statistics

#### request:

```json
{
  "id": "qwerty"
}
```

#### response:

```json
{
  "lobbies": [
    {
      "id": "string",
      "author_id": "string",
      "name": "string",
      "course_id": "string",
      "missions": [{
          "id": "string",
          "summary": "string",
          "users": [
            {
              "user_id": "string",
              "summary": "string",
              "total_diamonds": 0,
              "image_url": "string",
              "last_completed_level": 0
            }
          ]
      }]
    }
  ]
}
```
