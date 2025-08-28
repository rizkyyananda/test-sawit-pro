--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

-- Started on 2025-08-28 09:36:12

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4855 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 1981227)
-- Name: t_estate; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_estate (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    width bigint NOT NULL,
    length bigint NOT NULL,
    CONSTRAINT chk_t_estate_length CHECK ((length > 0)),
    CONSTRAINT chk_t_estate_width CHECK ((width > 0))
);


ALTER TABLE public.t_estate OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 1981235)
-- Name: t_tree; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.t_tree (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    estate_id uuid NOT NULL,
    x bigint NOT NULL,
    y bigint NOT NULL,
    height bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_t_tree_height CHECK (((height >= 1) AND (height <= 30))),
    CONSTRAINT chk_t_tree_x CHECK ((x > 0)),
    CONSTRAINT chk_t_tree_y CHECK ((y > 0))
);


ALTER TABLE public.t_tree OWNER TO postgres;

--
-- TOC entry 4848 (class 0 OID 1981227)
-- Dependencies: 215
-- Data for Name: t_estate; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_estate (id, width, length) FROM stdin;
754f298f-2a2f-44dd-b6b5-701ebac0a8f5	100	100
a6910cf9-326a-44a1-bd02-277b98362156	900	800
b8c0880e-ea4d-4ef1-b5b4-0b6265424722	50	50
cfb5040c-837d-4def-8880-0e9e22fbff0b	50000	50000
e45db355-e07c-42c0-9249-13d1404dbb47	50	50
187a398d-fb5f-424a-bed3-ffc0258e91db	50	50
20d6d707-d1cc-42d2-8719-9fe3f9bb9eb5	50	50
585d33dc-9a20-4111-b6e9-bc81dc0659a2	100	50
d125f09d-a9d4-4175-9e4f-d0c886f29f35	100	50
\.


--
-- TOC entry 4849 (class 0 OID 1981235)
-- Dependencies: 216
-- Data for Name: t_tree; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.t_tree (id, estate_id, x, y, height, created_at, updated_at, deleted_at) FROM stdin;
de2c6dc1-80da-4aea-9f8b-3ddc2dd207dd	754f298f-2a2f-44dd-b6b5-701ebac0a8f5	30	30	30	2025-07-05 13:14:35.128389+07	2025-07-05 13:14:35.128389+07	\N
fe21f1a1-004b-4d1e-be1e-1027b943f0b6	a6910cf9-326a-44a1-bd02-277b98362156	30	30	30	2025-07-06 22:32:59.427066+07	2025-07-06 22:32:59.427066+07	\N
2cb9e84f-f4b9-4aa4-86b2-38059beb4548	b8c0880e-ea4d-4ef1-b5b4-0b6265424722	10	10	25	2025-08-27 14:05:26.265056+07	2025-08-27 14:05:26.265056+07	\N
ab476895-208b-4fda-b5ed-0846b0025115	20d6d707-d1cc-42d2-8719-9fe3f9bb9eb5	10	10	25	2025-08-27 16:15:20.109704+07	2025-08-27 16:15:20.109704+07	\N
918f1643-6b1f-465c-a70b-0cbbe00c3154	585d33dc-9a20-4111-b6e9-bc81dc0659a2	10	10	30	2025-08-27 16:16:52.51941+07	2025-08-27 16:16:52.51941+07	\N
bed4673b-c7be-43fe-b02b-ae4e065bf898	d125f09d-a9d4-4175-9e4f-d0c886f29f35	10	10	30	2025-08-27 16:17:38.724814+07	2025-08-27 16:17:38.724814+07	\N
\.


--
-- TOC entry 4699 (class 2606 OID 1981234)
-- Name: t_estate t_estate_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_estate
    ADD CONSTRAINT t_estate_pkey PRIMARY KEY (id);


--
-- TOC entry 4703 (class 2606 OID 1981243)
-- Name: t_tree t_tree_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_tree
    ADD CONSTRAINT t_tree_pkey PRIMARY KEY (id);


--
-- TOC entry 4700 (class 1259 OID 1981249)
-- Name: idx_t_tree_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_t_tree_deleted_at ON public.t_tree USING btree (deleted_at);


--
-- TOC entry 4701 (class 1259 OID 1981250)
-- Name: idx_t_tree_estate_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_t_tree_estate_id ON public.t_tree USING btree (estate_id);


--
-- TOC entry 4704 (class 2606 OID 1981244)
-- Name: t_tree fk_t_estate_trees; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.t_tree
    ADD CONSTRAINT fk_t_estate_trees FOREIGN KEY (estate_id) REFERENCES public.t_estate(id) ON DELETE CASCADE;


-- Completed on 2025-08-28 09:36:13

--
-- PostgreSQL database dump complete
--

