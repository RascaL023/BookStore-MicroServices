package com.book.book_service.dto.response;

import java.time.LocalDateTime;
import java.util.List;


import lombok.Getter;
import lombok.Setter;

@Getter @Setter
public class BookResponse {
    Long id;
    String title;
    List<String> writers;
    Long price;

    LocalDateTime createdAt;
    LocalDateTime updatedAt;
    Boolean available;
}
