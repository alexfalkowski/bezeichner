# frozen_string_literal: true

Before('@pg') do
  Bezeichner.pg.create
end

After('@pg') do
  Bezeichner.pg.destroy
end
