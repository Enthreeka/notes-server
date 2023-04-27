

create table notes(
    id int generated always as identity,
    notes text,
    created_at timestamp DEFAULT now(),
    primary key (id)
);
