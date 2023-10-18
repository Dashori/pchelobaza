CREATE TABLE "bee_user_conf"(
    "id_user" INTEGER NOT NULL,
    "id_conf" INTEGER NOT NULL
);

ALTER TABLE
    "bee_user_conf" ADD PRIMARY KEY("id_user", "id_conf");

CREATE TABLE "bee_request"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "id_admin" INTEGER NOT NULL,
    "info" TEXT NOT NULL,
    "state" TEXT NOT NULL
);

ALTER TABLE
    "bee_request" ADD PRIMARY KEY("id");


CREATE TABLE "bee_farm_honey"(
    "id_farm" INTEGER NOT NULL,
    "id_honey" INTEGER NOT NULL
);
ALTER TABLE
    "bee_farm_honey" ADD PRIMARY KEY("id_farm", "id_honey");


CREATE TABLE "bee_comment"(
    "id" SERIAL NOT NULL,
    "id_conf" INTEGER NOT NULL,
    "id_user" INTEGER NOT NULL,
    "time" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "bee_comment" ADD PRIMARY KEY("id");


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
    "registered_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "role" TEXT NOT NULL
);
ALTER TABLE
    "bee_user" ADD PRIMARY KEY("id");


CREATE TABLE "bee_conf"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "data" TEXT NOT NULL,
    "place" TEXT NOT NULL,
    "capacity" INTEGER NOT NULL
);
ALTER TABLE
    "bee_conf" ADD PRIMARY KEY("id");


CREATE TABLE "bee_farm"(
    "id" SERIAL NOT NULL,
    "id_user" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "desrciption" TEXT NOT NULL,
    "address" TEXT NOT NULL
);
ALTER TABLE
    "bee_farm" ADD PRIMARY KEY("id");
ALTER TABLE
    "bee_farm_honey" ADD CONSTRAINT "bee_farm_honey_id_farm_foreign" FOREIGN KEY("id_farm") REFERENCES "bee_farm"("id");
ALTER TABLE
    "bee_farm_honey" ADD CONSTRAINT "bee_farm_honey_id_honey_foreign" FOREIGN KEY("id_honey") REFERENCES "bee_honey"("id");
ALTER TABLE
    "bee_user_conf" ADD CONSTRAINT "bee_user_conf_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_farm" ADD CONSTRAINT "bee_farm_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_comment" ADD CONSTRAINT "bee_comment_id_conf_foreign" FOREIGN KEY("id_conf") REFERENCES "bee_conf"("id");
ALTER TABLE
    "bee_conf" ADD CONSTRAINT "bee_conf_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_request" ADD CONSTRAINT "bee_request_id_user_foreign" FOREIGN KEY("id_user") REFERENCES "bee_user"("id");
ALTER TABLE
    "bee_user_conf" ADD CONSTRAINT "bee_user_conf_id_conf_foreign" FOREIGN KEY("id_conf") REFERENCES "bee_conf"("id");
ALTER TABLE
    "bee_request" ADD CONSTRAINT "bee_request_id_admin_foreign" FOREIGN KEY("id_admin") REFERENCES "bee_user"("id");


insert into bee_honey(name, description) values 
('Caldwell0','12345'),
('Miller1','12345'),
('Espinoza2','12345');