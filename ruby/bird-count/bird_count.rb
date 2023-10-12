class BirdCount
  @week = [0, 2, 5, 3, 7, 8, 4]

  def self.last_week
    @week
  end

  def initialize(birds_per_day)
    @week = birds_per_day
  end

  def yesterday
    @week[-2]
  end

  def total
    @week.sum
  end

  def busy_days
    @week.select {|day| day > 4}.size
  end

  def day_without_birds?
    @week.include? 0
  end
end
