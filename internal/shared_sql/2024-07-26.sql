/**
    An order can be of two types:
    - p2p: The buyer and seller will facilitate the transaction themselves
    - online: the order will be paid in advance and the seller will ship the product
 */
create type order_type as enum ('p2p', 'online');

/**
  An order starts its life in the 'created' state.
  It the order type is p2p, it instantly goes to the 'confirmed' state.
  If the order type is `online`, it goes to the 'confirmed' state after the user pays for it.
 */
create type order_state as enum ('created', 'confirmed', 'done', 'rejected');

alter table orders add column type order_type default 'p2p'::order_type not null;

alter table orders add column state order_state default 'created'::order_state not null;