package com.book.book_service.dto.request;

import java.util.ArrayList;
import java.util.List;

import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
public class BookRequest {
    @NotBlank(message = "Judul tidak boleh kosong")
    @Size(max = 30, message = "Judul terlalu panjang")
    String title;

    @NotEmpty(message = "ID penulis tidak boleh kosong")
    List<Long> writerIds = new ArrayList<>();

    @NotNull(message = "Harga tidak boleh kosong")
    @Min(value = 5000, message = "Harga minimal 5.000")
    Long price;
}
