class LogLineParser
  def initialize(line)
    tokens = line.split ':'

    @log_level = tokens[0][1..-2].downcase
    @message = tokens[1].lstrip.rstrip
  end

  def message
    @message
  end

  def log_level
    @log_level
  end

  def reformat
    "#{@message} (#{@log_level})"
  end
end
