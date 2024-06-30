defmodule Username do
  @moduledoc false

  @doc """
  Sanitize a username.
  """
  @spec sanitize(username:: charlist()) :: charlist()
  def sanitize([]), do: [] # Base case
  def sanitize([h|t]) do  # Recursive Case
    sanitized = case h do
      ?ä -> ~c"ae"
      ?ö -> ~c"oe"
      ?ü -> ~c"ue"
      ?ß -> ~c"ss"
      c when (c >= ?a and c <= ?z) -> [c]
      ?_ -> ~c"_"
      _ -> ~c""
    end

    sanitized ++ sanitize(t)
  end
end
