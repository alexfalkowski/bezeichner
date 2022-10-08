# frozen_string_literal: true

Bezeichner.pg.create

at_exit do
  Bezeichner.pg.destroy
end
