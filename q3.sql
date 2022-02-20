-- The book table contains a list of name, author, isbn number, description and the publish date.

--Exercise 1-1: Create a database named library
CREATE DATABASE library;

--Drops table books if it exists
DROP TABLE IF EXISTS books;

--Exercise 1-2: Create a table named books
CREATE TABLE books (
    book_id serial PRIMARY KEY,
    name text NOT NULL,
    author text NOT NULl,
    isbn VARCHAR (10) NOT NULL,
    description text NOT NULL,
    publication_date date NOT NULL DEFAULT NOW()
);


--Exercise 1-3: Populate the books table with three rows
INSERT INTO
    books (name, author, isbn, description, publication_date)
VALUES
    ('Introduction to the design & analysis of algorithms Edition: 3','Anany Levitin', '0132316811', 'This book is designed to teach techniques for the design and analysis of efficient computer algorithms through theoretical backgrounds and examples of advanced methods and data structures.', '2012-01-01'),
    ('Software Engineering Edition: 10','Ian Sommerville', '0133943038', ' Computer Science; Computers & Technology; New, Used & Rental Textbooks; Programming; Programming Languages; Software Design & Engineering; Software Design, Testing & Engineering; Software Development; Specialty Boutique', '2015-01-01'),
    ('Slackware Linux Essentials - The Official Guide to Slackware Linux Edition: 2','David Cantrell, Logan Johnson, Alan Hicks, Chris Lumens', '1571763384', 'Emacs, filesystem, GNU, HOWTO, installation, kernel, Linux, lilo, loadlin, network, partitioning, patches, process, root, security, setup, shell, Slackware, superuser, system configuration, vi, X-Windows', '2005-01-01' );
