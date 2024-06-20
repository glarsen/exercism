defmodule NameBadge do
  @spec print(number(), String.t(), String.t()) :: String.t()
  def print(id, name, department) do
    prefix = if id, do: "[#{id}] - ", else: ""
    suffix = if department, do: " - #{String.upcase(department)}", else: " - OWNER"

    prefix <> name <> suffix
  end
end
