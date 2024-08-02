--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7 (Debian 15.7-1.pgdg120+1)
-- Dumped by pg_dump version 16.3

-- Started on 2024-07-31 21:09:31

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

CREATE SCHEMA IF NOT EXISTS public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 3350 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 24577)
-- Name: messages; Type: TABLE; Schema: public; Owner: baseuser
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    title character varying(200) NOT NULL,
    data bytea,
    processed boolean DEFAULT false NOT NULL
);


ALTER TABLE public.messages OWNER TO baseuser;

--
-- TOC entry 214 (class 1259 OID 24576)
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: baseuser
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.messages_id_seq OWNER TO baseuser;

--
-- TOC entry 3351 (class 0 OID 0)
-- Dependencies: 214
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baseuser
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- TOC entry 3199 (class 2604 OID 24580)
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- TOC entry 3202 (class 2606 OID 24585)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


-- Completed on 2024-07-31 21:09:32

--
-- PostgreSQL database dump complete
--

