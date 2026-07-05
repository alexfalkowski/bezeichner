# frozen_string_literal: true

Then('I should receive a not found error from gRPC') do
  expect(@response).to be_a(GRPC::NotFound)
end

Then('I should receive an internal error from gRPC') do
  expect(@response).to be_a(GRPC::Internal)
end

Then('I should receive an invalid argument error from gRPC') do
  expect(@response).to be_a(GRPC::InvalidArgument)
end
