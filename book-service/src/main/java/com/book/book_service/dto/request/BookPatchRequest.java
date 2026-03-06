package com.book.book_service.dto.request;

import java.util.List;

import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.Size;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
public class BookPatchRequest {
    @Size(max = 30, message = "Judul terlalu panjang")
    String title;

    List<Long> writerIds;

    @Min(value = 5000, message = "Harga minimal 5.000")
    Long price;
}
