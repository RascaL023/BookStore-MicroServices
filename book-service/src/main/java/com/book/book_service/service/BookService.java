package com.book.book_service.service;

import java.time.LocalDateTime;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import com.book.book_service.dto.mapper.BookMapper;
import com.book.book_service.dto.request.BookPatchRequest;
import com.book.book_service.dto.request.BookRequest;
import com.book.book_service.dto.response.BookResponse;
import com.book.book_service.model.Book;
import com.book.book_service.repository.BookRepository;


@Service
public class BookService {
    @Autowired private BookRepository bookRepository;

    private Book getBook(Long id) {
        return bookRepository.findById(id)
            .orElseThrow(() -> new RuntimeException(
                "Buku tidak ditemukkan"
            )
        );
    }


    public BookResponse getById(Long id) {
        return BookMapper.toResponse(
            getBook(id)
        );
    }

    public BookResponse upsert(
        Long id, Object request
    ) {
        Book book;
        if (id == null) {
            book = BookMapper.toEntity(null, (BookRequest)request);
            book.setCreatedAt(LocalDateTime.now());
            book.setUpdatedAt(LocalDateTime.now());
            book.setAvailable(true);
            
            // !!!writer validation
            return BookMapper.toResponse(
                bookRepository.save(book)
            );
        } 
        
        book = getBook(id);
        if (request instanceof BookRequest putRequest) {
            book = BookMapper.toEntity(book, putRequest);
            // !!!writer validation
        } else {
            BookPatchRequest patchRequest = 
                (BookPatchRequest)request;
            book = BookMapper.toEntity(book, patchRequest);
            if (patchRequest.getWriterIds() != null) {
                // !!!writer validation
            }
        }

        book.setUpdatedAt(LocalDateTime.now());
        return BookMapper.toResponse(
            bookRepository.save(book)
        );
    }

    public BookResponse deleteById(Long id) {
        Book book = getBook(id);
        book.setAvailable(false);
        return BookMapper.toResponse(
            bookRepository.save(book)
        );
    }


    public Page<BookResponse> getAll(Pageable pageable) {
        Page<Book> page = bookRepository.findAll(pageable);
        List<BookResponse> content = page.getContent().stream()
            .map(BookMapper::toResponse).toList();

        return new PageImpl<>(
            content, 
            pageable, 
            page.getTotalElements()
        );
    }
}
