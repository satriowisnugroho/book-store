--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Debian 14.1-1.pgdg110+1)
-- Dumped by pg_dump version 14.11 (Homebrew)

--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: root
--
TRUNCATE public.books RESTART IDENTITY CASCADE;

COPY public.books (isbn, title, price, created_at, updated_at) FROM stdin;
lorem	ipsum	123	2024-06-06 15:37:14	2024-06-06 15:37:14
ipsum	lorem	123	2024-06-06 15:37:14	2024-06-06 15:37:14
\.

--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: root
--

TRUNCATE public.orders RESTART IDENTITY CASCADE;

COPY public.orders (user_id, book_id, quantity, price, fee, total_price, created_at, updated_at) FROM stdin;
1	1	4	25000	1000	101000	2024-06-07 22:43:27.750419	2024-06-07 22:43:27.750419
2	2	4	20000	1000	81000	2024-06-07 22:43:27.750419	2024-06-07 22:43:27.750419
\.

--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

TRUNCATE public.users RESTART IDENTITY CASCADE;

COPY public.users (email, fullname, crypted_password, created_at, updated_at) FROM stdin;
foo@gmail.com	bar	qwerty	2024-06-06 15:37:14	2024-06-06 15:37:14
bar@gmail.com	foo	qwerty	2024-06-06 15:37:14	2024-06-06 15:37:14
\.

--
-- PostgreSQL database dump complete
--
