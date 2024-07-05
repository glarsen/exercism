defmodule PaintByNumber do
  @spec palette_bit_size(non_neg_integer()) :: non_neg_integer()
  def palette_bit_size(color_count), do: palette_bit_size(color_count, 0)

  defp palette_bit_size(color_count, exponent) do
    if Integer.pow(2, exponent) >= color_count do
      exponent
    else
      palette_bit_size(color_count, exponent + 1)
    end
  end

  @spec empty_picture() :: <<>>
  def empty_picture() do
    <<>>
  end

  @spec test_picture() :: <<_::8>>
  def test_picture() do
    <<0::2, 1::2, 2::2, 3::2>>
  end

  def prepend_pixel(picture, color_count, pixel_color_index) do
    <<pixel_color_index::size(palette_bit_size(color_count)), picture::bitstring>>
  end

  @spec get_first_pixel(bitstring(), non_neg_integer()) :: bitstring()
  def get_first_pixel(<<>>, _), do: nil

  def get_first_pixel(picture, color_count) do
    bits = palette_bit_size(color_count)
    <<first::size(bits), _::bitstring>> = picture
    first
  end

  @spec drop_first_pixel(bitstring(), non_neg_integer()) :: bitstring()
  def drop_first_pixel(<<>>, _), do: <<>>

  def drop_first_pixel(picture, color_count) do
    bits = palette_bit_size(color_count)
    <<_::size(bits), rest::bitstring>> = picture
    rest
  end

  @spec concat_pictures(bitstring(), bitstring()) :: bitstring()
  def concat_pictures(picture1, picture2) do
    <<picture1::bitstring, picture2::bitstring>>
  end
end
