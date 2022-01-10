/* create table of accessory */
CREATE TABLE IF NOT EXISTS artifact_accessory (
    id SERIAL PRIMARY KEY NOT NULL,
    /*
       the artifact id of the accessory itself.
    */
    artifact_id bigint,
    /*
     the subject artifact id of the accessory.
    */
    subject_artifact_id bigint,
    /*
     the type of the accessory, like signature.cosign.
    */
    type varchar(256),
    size bigint,
    digest varchar(1024),
    creation_time timestamp default CURRENT_TIMESTAMP,
    FOREIGN KEY (artifact_id) REFERENCES artifact(id),
    FOREIGN KEY (subject_artifact_id) REFERENCES artifact(id),
    CONSTRAINT unique_artifact_accessory UNIQUE (artifact_id, subject_artifact_id)
);

ALTER TABLE notification_policy ADD COLUMN IF NOT EXISTS repository varchar(256);
