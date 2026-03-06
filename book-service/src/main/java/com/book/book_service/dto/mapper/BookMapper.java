package com.book.book_service.dto.mapper;

import com.book.book_service.dto.request.BookPatchRequest;
import com.book.book_service.dto.request.BookRequest;
import com.book.book_service.dto.response.BookResponse;
import com.book.book_service.model.Book;

public class BookMapper {
    public static Book toEntity(Book book, BookRequest request) {
        if (book == null) book = new Book();
        book.setTitle(request.getTitle());
        book.setPrice(request.getPrice());

        return book;
    }

    public static Book toEntity(Book book, BookPatchRequest request) {
        if (book == null) book = new Book();
        if (request.getTitle() != null)
            book.setTitle(request.getTitle());

        if (request.getPrice() != null)
            book.setPrice(request.getPrice());

        return book;
    }

    public static BookRequest toRequest(Book book) {
        BookRequest request = new BookRequest();
        request.setTitle(book.getTitle());
        request.setPrice(book.getPrice());
        request.setWriterIds(book.getWriterIds());

        return request;
    }

    public static BookPatchRequest toPatchRequest(Book book) {
        BookPatchRequest request = new BookPatchRequest();
        request.setTitle(book.getTitle());
        request.setPrice(book.getPrice());
        request.setWriterIds(book.getWriterIds());

        return request;
    }

    public static BookResponse toResponse(Book book) {
        BookResponse response = new BookResponse();
        response.setId(book.getId());
        response.setTitle(book.getTitle());
        response.setPrice(book.getPrice());
        response.setAvailable(book.getAvailable());
        response.setCreatedAt(book.getCreatedAt());
        response.setUpdatedAt(book.getUpdatedAt());

        return response;
    }
}
