BEGIN;

SET client_encoding = 'UTF8';

CREATE TABLE facts (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT,
    links       TEXT
);

INSERT INTO facts(title, description, links)
VALUES('високосная секунда', 'Дополнительная (висоќосная, скачущая) сеќунда[4][5][6], или сеќунда координ́ации[7] (англ. leap second) — секунда, иногда добавляемая (теоретически возможно и вычитание) в шкалу всемирного координированного времени (UTC) для согласования его со средним солнечным временем UT1.', 'shorturl.at/dwzK9');

COMMIT;
