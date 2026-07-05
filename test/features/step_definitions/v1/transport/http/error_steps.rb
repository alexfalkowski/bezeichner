# frozen_string_literal: true

Then('I should receive a not found error from HTTP') do
  expect(@response.code).to eq(404)
end

Then('I should receive an internal error from HTTP') do
  expect(@response.code).to eq(500)
end

Then('I should receive an invalid argument error from HTTP') do
  expect(@response.code).to eq(400)
end
