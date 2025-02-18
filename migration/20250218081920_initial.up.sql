create table companies
(
    id         integer
        primary key autoincrement,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    code       text,
    member_id  text,
    name       text,
    address    text,
    email      text,
    website    text,
    "index"    text,
    sector     text,
    market     text,
    icon       text
);

create unique index idx_companies_code
    on companies (code);

create index idx_companies_deleted_at
    on companies (deleted_at);

--bun:split

create table company_shares
(
    id         integer
        primary key autoincrement,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    company_id integer
        constraint fk_companies_shares
            references companies,
    date       datetime
);

create index idx_company_shares_date
    on company_shares (date);

create index idx_company_shares_deleted_at
    on company_shares (deleted_at);

--bun:split

create table company_share_holders
(
    company_share_id  integer
        constraint fk_company_shares_share_holders
            references company_shares
            on update cascade on delete cascade,
    title             text,
    capital_by_amount real,
    capital_by_volume real,
    vote_right        real
);
