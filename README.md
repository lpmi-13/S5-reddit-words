# S5-reddit-words

I'm planning to compile a corpus of programming text, and various programming subreddits are going to be a large component of that.

This service will take a subreddit, or group of subreddits and grab the top 30 newest posts' comments, then put them in a datastore somewhere (probably mongoDB, but maybe eventually S3).

If I feel frisky, I might even start instrumenting code and pushing metrics to prometheus pushgateway.
