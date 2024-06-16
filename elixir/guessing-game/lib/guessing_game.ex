defmodule GuessingGame do
  def compare(secret, guess \\ :no_guess)
  def compare(_, :no_guess), do: "Make a guess"
  def compare(secret, guess) when secret == guess, do: "Correct"
  def compare(secret, guess) when abs(secret - guess) == 1, do: "So close"
  def compare(secret, guess) when guess > secret, do: "Too high"
  def compare(secret, guess) when guess < secret, do: "Too low"
end
