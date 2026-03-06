package com.book.book_service.controller.api;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.book.book_service.dto.request.BookRequest;
import com.book.book_service.service.BookService;

import jakarta.validation.Valid;

@RestController
@RequestMapping("/api/books")
public class BookApiController {
    @Autowired private BookService bookService;


    @PostMapping
    public ResponseEntity<?> add(
        @Valid @RequestBody BookRequest request
    ) {
        return ResponseEntity.ok(
            bookService.upsert(null, request)
        );
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> getById(
        @PathVariable("id") Long id
    ) {
        return ResponseEntity.ok(
            bookService.getById(id)
        );
    }
}
