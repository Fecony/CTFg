<script lang="ts">
  import { Button, Heading, P } from 'flowbite-svelte';
  const initialFlag = 'flag{you_found_twitter}';
</script>

<div class="text-center">
  <div class="mb-4">
    <Heading tag="h3">Welcome to CTFg!</Heading>
    <P class="text-center">Register/Login to get started.</P>
  </div>
  <hr />
  <P>Are you ready to embark on a thrilling journey into the world of cybersecurity and digital forensics? Today, you will step into the shoes of a detective and use your skills to solve a complex cybercrime that has left law enforcement stumped.</P><br />
  <P>A murder has taken place, and it's up to you to find the killer. You'll have to analyze digital evidence, comb through logs, and even try your hand at hacking to uncover the truth. As you dive deeper into the investigation, you'll encounter obstacles and challenges that will push your skills and abilities to the limit. You'll need to think outside the box and use your creativity to overcome these hurdles and bring the perpetrator to justice.  Finally, write up your evidence and conclusions so they can stand up in "court."</P><br />
  <P>The stakes are high, and failure is not an option. The fate of the case rests in your hands, and the pressure is on to solve the crime before it's too late. Are you up for the challenge? Can you rise to the occasion and bring the killer to justice? Login to <a style="color: blue" href="http://twitter-flask.chals.mcpshsf.com/">http://twitter-flask.chals.mcpshsf.com/</a> with the username <b>“sadamana”</b> and password <b>“s4d4m4n4”</b> to get started.</P><br />
  <hr />
  <div class="mt-4">
    <P class="text-center">For help related to the competition, go to the wiki. It will help you a lot.</P>
    <Button color="blue" href="https://wiki.ctfg.io">wiki</Button>
  </div>
</div>
