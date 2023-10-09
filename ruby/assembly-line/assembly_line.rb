class AssemblyLine
  HOURLY_RATE = 221.freeze

  def initialize(speed)
    raise ArgumentError, 'speed must be between 1 and 10' unless speed.between? 1, 10
    @speed = speed
  end

  def production_rate_per_hour
    case @speed
    when 0
      0.0
    when 1..4
      @speed * HOURLY_RATE.to_f
    when 5..8
      @speed * HOURLY_RATE * 0.9
    when 9
      @speed * HOURLY_RATE * 0.8
    when 10
      @speed * HOURLY_RATE * 0.77
    end
  end

  def working_items_per_minute
    self.production_rate_per_hour.div(60)
  end
end
