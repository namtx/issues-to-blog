
---
label: til
layout: default
title: aggregate_failures in RSpec
---
`aggregate_failues` is an API of `RSpec::Expectations`, allows you to group a set of all failures at once, rather than its aborting when see first failure.

```ruby
describe '.from_string', :aggregate_failures do
  it 'something failure' do
  end

  it 'also failure too' do
  end
end
```

