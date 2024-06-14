defmodule FreelancerRates do
  def daily_rate(hourly_rate), do: hourly_rate * 8.0

  defp monthly_rate(hourly_rate), do: daily_rate(hourly_rate) * 22

  def apply_discount(before_discount, discount), do: before_discount - (before_discount * (discount / 100.0))

  def monthly_rate(hourly_rate, discount) do
    hourly_rate |> monthly_rate |> apply_discount(discount) |> ceil
  end

  def days_in_budget(budget, hourly_rate, discount) do
    Float.floor(budget / apply_discount(daily_rate(hourly_rate), discount), 1)
  end
end
