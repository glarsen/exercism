defmodule Lasagna do
  @moduledoc """
  Bake a lasagna!
  """

  @spec expected_minutes_in_oven() :: 40
  def expected_minutes_in_oven, do: 40

  @spec remaining_minutes_in_oven(number()) :: number()
  def remaining_minutes_in_oven(t), do: expected_minutes_in_oven() - t

  @spec preparation_time_in_minutes(number()) :: number()
  def preparation_time_in_minutes(layers), do: layers * 2

  @spec total_time_in_minutes(number(), number()) :: number()
  def total_time_in_minutes(layers, time), do: preparation_time_in_minutes(layers) + time

  @spec alarm() :: String.t()
  def alarm, do: "Ding!"
end
