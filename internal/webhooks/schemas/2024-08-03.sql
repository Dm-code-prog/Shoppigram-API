create type payment_providers as enum ('cloud_payments');

create type payments_event_type as enum ('check', 'pay');

create table payments_extra_info
(
    id                   uuid primary key default uuid_generate_v4(),
    created_at           timestamp        default now(),
    updated_at           timestamp        default now(),
    order_id             uuid references orders (id),
    provider             payment_providers   not null,
    order_state_snapshot order_state         not null,
    event_type           payments_event_type not null,
    extra_info           json                not null
);