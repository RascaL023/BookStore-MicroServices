package com.book.book_service.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.book.book_service.model.Book;

public interface BookRepository extends JpaRepository<Book, Long> {
    
}
