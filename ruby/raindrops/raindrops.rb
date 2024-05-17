class Raindrops
  def self.convert(number)
    sounds = {
      3 => "Pling",
      5 => "Plang",
      7 => "Plong"
    }

    result = sounds
               .select { |k, _| number % k == 0 }
               .values
               .join

    result.empty? ? number.to_s : result
  end
end