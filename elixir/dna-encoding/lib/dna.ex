defmodule DNA do
  def encode_nucleotide(code_point) do
    case code_point do
      ?\s -> 0b0000
      ?A -> 0b0001
      ?C -> 0b0010
      ?G -> 0b0100
      ?T -> 0b1000
    end
  end

  def decode_nucleotide(encoded_code) do
    case encoded_code do
      0b0000 -> ?\s
      0b0001 -> ?A
      0b0010 -> ?C
      0b0100 -> ?G
      0b1000 -> ?T
    end
  end

  @spec encode(charlist()) :: bitstring()
  def encode(dna), do: do_encode(dna, <<>>)

  @spec do_encode(charlist(), bitstring()) :: bitstring()
  defp do_encode([], acc), do: acc

  defp do_encode([nucleotide | remaining], acc) do
    do_encode(remaining, <<acc::bitstring, encode_nucleotide(nucleotide)::4>>)
  end

  @spec decode(bitstring()) :: charlist()
  def decode(dna), do: do_decode(dna, [])

  @spec do_decode(bitstring(), charlist()) :: charlist()
  defp do_decode(<<>>, acc), do: acc

  defp do_decode(<<nucleotide::4, remaining::bitstring>>, acc) do
    do_decode(remaining, acc ++ [decode_nucleotide(nucleotide)])
  end
end
