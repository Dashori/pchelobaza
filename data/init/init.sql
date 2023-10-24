CREATE TABLE "bee_user_conference"(
    "id_user" INTEGER NOT NULL,
    "id_conference" INTEGER NOT NULL
);
ALTER TABLE
    "bee_user_conference" ADD PRIMARY KEY("id_user", "id_conference");

CREATE TABLE "bee_request"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "status" TEXT NOT NULL
);
ALTER TABLE
    "bee_request" ADD PRIMARY KEY("id");
ALTER TABLE
    "bee_request" ADD CONSTRAINT "bee_request_id_user_unique" UNIQUE(id_user);
ALTER TABLE
    "bee_request"  ADD CONSTRAINT "status_type_check" CHECK (status IN ('approve', 'rejected', 'waiting'));


CREATE TABLE "bee_farm_honey"(
    "id_farm" INTEGER NOT NULL,
    "id_honey" INTEGER NOT NULL
);
ALTER TABLE
    "bee_farm_honey" ADD PRIMARY KEY("id_farm", "id_honey");

    
CREATE TABLE "bee_review"(
    "id" SERIAL NOT NULL,
    "id_conference" INTEGER NOT NULL,
    "id_user" INTEGER NOT NULL,
    "date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "bee_review" ADD PRIMARY KEY("id");

CREATE TABLE "bee_honey"(
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "bee_honey" ADD PRIMARY KEY("id");

CREATE TABLE "bee_user"(
    "id" SERIAL NOT NULL,
    "login" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "surname" TEXT NOT NULL,
    "contact" TEXT NOT NULL,
    "registered_at" DATE NOT NULL,
    "role" TEXT NOT NULL
);
ALTER TABLE
    "bee_user" ADD PRIMARY KEY("id");
ALTER TABLE
    "bee_user" ADD CONSTRAINT "bee_user_login_unique" UNIQUE(login);
ALTER TABLE
    "bee_user"  ADD CONSTRAINT "role_type_check" CHECK (role IN ('beeadmin', 'beeman', 'beemaster'));


CREATE TABLE "bee_conference"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "date" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "address" TEXT NOT NULL,
    "maximum_users" INTEGER NOT NULL,
    "current_users" INTEGER NOT NULL
);
ALTER TABLE
    "bee_conference" ADD PRIMARY KEY("id");
ALTER TABLE
    "bee_conference" ADD CONSTRAINT "bee_user_name_unique" UNIQUE(name);

CREATE TABLE "bee_farm"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "address" TEXT NOT NULL
);
ALTER TABLE
    "bee_farm" ADD PRIMARY KEY("id");
ALTER TABLE
    "bee_farm" ADD CONSTRAINT "bee_farm_login_unique" UNIQUE(name);

ALTER TABLE
    "bee_farm_honey" ADD CONSTRAINT "bee_farm_honey_id_farm_foreign" FOREIGN KEY("id_farm") REFERENCES "bee_farm"("id");
ALTER TABLE
    "bee_farm_honey" ADD CONSTRAINT "bee_farm_honey_id_honey_foreign" FOREIGN KEY("id_honey") REFERENCES "bee_honey"("id");

ALTER TABLE
    "bee_user_conference" ADD CONSTRAINT "bee_user_conference_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_farm" ADD CONSTRAINT "bee_farm_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_review" ADD CONSTRAINT "bee_review_id_conference_foreign" FOREIGN KEY("id_conference") REFERENCES "bee_conference"("id");
ALTER TABLE
    "bee_conference" ADD CONSTRAINT "bee_conference_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_request" ADD CONSTRAINT "bee_request_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_user_conference" ADD CONSTRAINT "bee_user_conference_id_conference_foreign" FOREIGN KEY("id_conference") REFERENCES "bee_conference"("id");