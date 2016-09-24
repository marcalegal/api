CREATE TABLE "users" (
  "id" serial NOT NULL UNIQUE,
  "email" varchar(50) NOT NULL,
  "password" varchar(128) NOT NULL,
  "name" varchar(255) NOT NULL,
  "lastname" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "kind" int NOT NULL DEFAULT '0',
  "session_token" varchar(128) NOT NULL DEFAULT '0',
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "deleted_at" timestamp NOT NULL,
  CONSTRAINT users_pk PRIMARY KEY ("email")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "brands" (
  "id" serial NOT NULL,
  "user_id" bigint NOT NULL,
  "n_request" varchar(10) NOT NULL DEFAULT 'XXXXXXX',
  "register_kind" varchar(255) NOT NULL,
  "name" varchar(50) NOT NULL,
  "kind" varchar(50) NOT NULL,
  "logo" TEXT NOT NULL,
  "logo description" TEXT NOT NULL,
  "conflict" BOOLEAN NOT NULL,
  "resume" TEXT NOT NULL,
  "total" DECIMAL NOT NULL,
  "dom_registry" BOOLEAN NOT NULL,
  "state" int NOT NULL DEFAULT '0',
  "attorney_power" TEXT NOT NULL,
  "payment_code" TEXT NOT NULL,
  CONSTRAINT brands_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "naturals" (
  "id" serial NOT NULL,
  "name" varchar(255) NOT NULL,
  "lastname" varchar(255) NOT NULL,
  "rut" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "comuna" varchar(255) NOT NULL,
  "city" varchar(255) NOT NULL,
  "country" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "belong_to_brand" bigint NOT NULL,
  CONSTRAINT naturals_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "juridicas" (
  "id" serial NOT NULL,
  "nombre_rpl" varchar(255) NOT NULL,
  "rut_rpl" varchar(255) NOT NULL,
  "nombre" varchar(255) NOT NULL,
  "razon" varchar(255) NOT NULL,
  "rut" varchar(255) NOT NULL,
  "giro" varchar(255) NOT NULL,
  "address" varchar(255) NOT NULL,
  "comuna" varchar(255) NOT NULL,
  "city" varchar(255) NOT NULL,
  "country" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL,
  "belong_to_brand" bigint NOT NULL,
  CONSTRAINT juridicas_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "values" (
  "id" serial NOT NULL,
  "tipo" varchar(50) NOT NULL,
  "monto" varchar(50) NOT NULL,
  "tipo_tarifa" varchar(50) NOT NULL,
  "update_time" varchar(50) NOT NULL,
  CONSTRAINT values_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "words" (
  "id" serial NOT NULL,
  "word" varchar(25) NOT NULL,
  CONSTRAINT words_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "clases" (
  "id" serial NOT NULL,
  "tipo" varchar(25) NOT NULL,
  "id_clase" bigint NOT NULL,
  "id_sub_clases" varchar(25) NOT NULL,
  "detalle" TEXT NOT NULL,
  CONSTRAINT clases_pk PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);




ALTER TABLE "brands" ADD CONSTRAINT "brands_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id");

ALTER TABLE "naturals" ADD CONSTRAINT "natural_fk0" FOREIGN KEY ("belong_to_brand") REFERENCES "brands"("id");

ALTER TABLE "juridicas" ADD CONSTRAINT "juridicas_fk0" FOREIGN KEY ("belong_to_brand") REFERENCES "brands"("id");
