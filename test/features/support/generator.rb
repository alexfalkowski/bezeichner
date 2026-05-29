# frozen_string_literal: true

%w[pg pg_alias].each do |name|
  Bezeichner.pg.create(name)
end

at_exit do
  %w[pg_alias pg].each do |name|
    Bezeichner.pg.destroy(name)
  end
end
