CREATE TABLE armies (
    "id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "name"  VARCHAR(255) NOT NULL,
    "cohort" json NOT NULL,
    "auxiliary" json NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT "armies_id" PRIMARY KEY ("id"),
    CONSTRAINT "armies_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO "armies" ("id", "user_id", "name", "cohort", "auxiliary", "created_at", "updated_at")
VALUES ('fabc3e24-5dca-47f4-86ba-b3e504de4ccb', '0dde213a-a81b-4b02-a665-41ec6c037112', 'Test Army', '{
    "units": {
        "15": {
            "templateId": "W_HORSE",
            "name": "Horse",
            "cost": 250,
            "color": "blue",
            "strength": 2,
            "health": 3,
            "speed": 3,
            "moxie": 45,
            "attackRange": 1,
            "attackType": "melee",
            "moveType": "basic",
            "onAttack": [],
            "onDie": [],
            "onKill": [],
            "onMove": [],
            "onStrike": [],
            "onStruck": [],
            "onTurnEnd": [],
            "activeAbilities": []
        }
    }
}', '[
    {
        "templateId": "W_HORSE",
        "name": "Horse",
        "cost": 250,
        "color": "blue",
        "strength": 2,
        "health": 3,
        "speed": 3,
        "moxie": 45,
        "attackRange": 1,
        "attackType": "melee",
        "moveType": "basic",
        "onAttack": [],
        "onDie": [],
        "onKill": [],
        "onMove": [],
        "onStrike": [],
        "onStruck": [],
        "onTurnEnd": [],
        "activeAbilities": []
    }
]', now(), now());

INSERT INTO "armies" ("id", "user_id", "name", "cohort", "auxiliary", "created_at", "updated_at")
VALUES ('abcd661c-18c4-4c7a-bd9d-e35ac06a48f5', '0c79665d-0ff8-4df9-8d9e-fe44b4b36308', 'Test Army', '{
    "units": {
        "15": {
            "templateId": "W_HORSE",
            "name": "Horse",
            "cost": 250,
            "color": "blue",
            "strength": 2,
            "health": 3,
            "speed": 3,
            "moxie": 45,
            "attackRange": 1,
            "attackType": "melee",
            "moveType": "basic",
            "onAttack": [],
            "onDie": [],
            "onKill": [],
            "onMove": [],
            "onStrike": [],
            "onStruck": [],
            "onTurnEnd": [],
            "activeAbilities": []
        }
    }
}', '[
    {
        "templateId": "W_HORSE",
        "name": "Horse",
        "cost": 250,
        "color": "blue",
        "strength": 2,
        "health": 3,
        "speed": 3,
        "moxie": 45,
        "attackRange": 1,
        "attackType": "melee",
        "moveType": "basic",
        "onAttack": [],
        "onDie": [],
        "onKill": [],
        "onMove": [],
        "onStrike": [],
        "onStruck": [],
        "onTurnEnd": [],
        "activeAbilities": []
    }
]', now(), now());
